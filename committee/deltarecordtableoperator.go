// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package committee

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"

	// require sqlite3 driver
	"github.com/pkg/errors"
	_ "modernc.org/sqlite"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-election/util"
)

// InsertDeltaRecordsFunc defines an api to insert records
type InsertDeltaRecordsFunc func(string, DRIVERTYPE, interface{}, *sql.Tx) (
	map[hash.Hash256]int,
	map[int64]hash.Hash256,
	error,
)

type deltaRecordTableOperator struct {
	recordTableOperator
	insertDeltaRecordsFunc InsertDeltaRecordsFunc
	indexQuery             string
}

// NewDeltaRecordTableOperator creates a record table storing delta
func NewDeltaRecordTableOperator(
	tableName string,
	driverName DRIVERTYPE,
	insertDeltaRecordsFunc InsertDeltaRecordsFunc,
	queryRecordsFunc QueryRecordsFunc,
	recordTableCreation string,
) (Operator, error) {
	var insertHeightToRecordsQuery, insertIdenticalQuery string
	switch driverName {
	case SQLITE:
		insertHeightToRecordsQuery = fmt.Sprintf("INSERT OR REPLACE INTO height_to_%s (height, ids, frequencies, indexes) VALUES (?, ?, ?, ?)", tableName)
		insertIdenticalQuery = fmt.Sprintf("INSERT OR IGNORE INTO identical_%s (height, identical_to) VALUES (?, ?)", tableName)
	case MYSQL:
		insertHeightToRecordsQuery = fmt.Sprintf("REPLACE INTO height_to_%s (height, ids, frequencies, indexes) VALUES (?, ?, ?, ?)", tableName)
		insertIdenticalQuery = fmt.Sprintf("INSERT IGNORE INTO identical_%s (height, identical_to) VALUES (?, ?)", tableName)
	default:
		return nil, errors.New("Wrong driver type")
	}
	return &deltaRecordTableOperator{
		recordTableOperator{
			tableName:                  tableName,
			driverName:                 driverName,
			frequencyQuery:             fmt.Sprintf("SELECT ids, frequencies FROM height_to_%s WHERE height = ?", tableName),
			hashQuery:                  fmt.Sprintf("SELECT id, hash FROM %s WHERE id IN (%s)", tableName, "%s"),
			idQuery:                    fmt.Sprintf("SELECT id, hash FROM %s WHERE hash IN ('%s')", tableName, "%s"),
			identicalQuery:             fmt.Sprintf("SELECT identical_to FROM identical_%s WHERE height = ?", tableName),
			lastHeightQuery:            fmt.Sprintf("SELECT MAX(max_height) FROM (SELECT MAX(height) AS max_height FROM identical_%s WHERE height < ? UNION SELECT MAX(height) AS max_height FROM height_to_%s WHERE height < ?) AS height", tableName, tableName),
			insertHeightToRecordsQuery: insertHeightToRecordsQuery,
			insertIdenticalQuery:       insertIdenticalQuery,
			queryRecordsFunc:           queryRecordsFunc,
			tableCreations: []string{
				fmt.Sprintf(recordTableCreation, tableName),
				fmt.Sprintf("CREATE TABLE IF NOT EXISTS height_to_%s (height INTEGER PRIMARY KEY, ids BLOB, frequencies BLOB, indexes BLOB)", tableName),
				fmt.Sprintf("CREATE TABLE IF NOT EXISTS identical_%s (height INTEGER PRIMARY KEY, identical_to INTEGER)", tableName),
			},
		},
		insertDeltaRecordsFunc,
		fmt.Sprintf("SELECT ids, indexes FROM height_to_%s WHERE height = ?", tableName),
	}, nil
}

func (op *deltaRecordTableOperator) TipHeight(db *sql.DB, tx *sql.Tx) (uint64, error) {
	return op.lastHeight(uint64(math.MaxInt64), db, tx)
}

func (op *deltaRecordTableOperator) Get(height uint64, db *sql.DB, tx *sql.Tx) (interface{}, error) {
	return op.recordTableOperator.Get(height, db, tx)
}

func (op *deltaRecordTableOperator) indexes(height uint64, sdb *sql.DB, tx *sql.Tx) (
	uint64,
	map[int64]hash.Hash256,
	error,
) {
	height, err := op.identicalTo(height, sdb, tx)
	if err != nil {
		return 0, nil, err
	}
	var (
		idBytes, idxBytes []byte
		ids               []int64
		index2ids         map[int64]int64
	)
	if tx != nil {
		err = tx.QueryRow(op.indexQuery, util.Uint64ToInt64(height)).Scan(&idBytes, &idxBytes)
	} else {
		err = sdb.QueryRow(op.indexQuery, util.Uint64ToInt64(height)).Scan(&idBytes, &idxBytes)
	}
	switch err {
	case sql.ErrNoRows:
		return height, nil, nil
	case nil:
	default:
		return 0, nil, err
	}
	if err = json.Unmarshal(idBytes, &ids); err != nil {
		return 0, nil, err
	}
	if err = json.Unmarshal(idxBytes, &index2ids); err != nil {
		return 0, nil, err
	}
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.Query(fmt.Sprintf(op.hashQuery, atos(ids)))
	} else {
		rows, err = sdb.Query(fmt.Sprintf(op.hashQuery, atos(ids)))
	}
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	id2hashes := make(map[int64]hash.Hash256, len(ids))
	indexes := make(map[int64]hash.Hash256, len(index2ids))
	for rows.Next() {
		var id int64
		var val string
		if err := rows.Scan(&id, &val); err != nil {
			return 0, nil, err
		}
		h, err := hash.HexStringToHash256(val)
		if err != nil {
			return 0, nil, err
		}
		id2hashes[id] = h
	}
	if rows.Err() != nil {
		return 0, nil, rows.Err()
	}
	for idx, id := range index2ids {
		h, ok := id2hashes[id]
		if !ok || h == hash.ZeroHash256 {
			return 0, nil, errors.Errorf("no record for id %d", id)
		}
		indexes[idx] = h
	}
	return height, indexes, nil
}

func (op *deltaRecordTableOperator) Put(height uint64, records interface{}, tx *sql.Tx) (err error) {
	var hash2Frequencies map[hash.Hash256]int
	var index2Hash map[int64]hash.Hash256
	var lastHeight uint64
	if hash2Frequencies, index2Hash, err = op.insertDeltaRecordsFunc(op.tableName, op.driverName, records, tx); err != nil {
		return err
	}
	if lastHeight, err = op.lastHeight(height, nil, tx); err != nil {
		return err
	}
	lastIdenticalHeight, lastIndexes, err := op.indexes(lastHeight, nil, tx)
	if err != nil {
		return errors.Wrap(err, "failed to get record hashes")
	}
	if len(hash2Frequencies) == 0 {
		if _, err := tx.Exec(op.insertIdenticalQuery, height, lastIdenticalHeight); err != nil {
			return err
		}
	} else {
		recordHashes := make([]string, 0, len(hash2Frequencies))
		for index, hash := range lastIndexes {
			if _, ok := index2Hash[index]; !ok {
				index2Hash[index] = hash
				if _, ok := hash2Frequencies[hash]; !ok {
					hash2Frequencies[hash] = 1
				} else {
					hash2Frequencies[hash]++
				}
			}
		}
		for h := range hash2Frequencies {
			recordHashes = append(recordHashes, hex.EncodeToString(h[:]))
		}
		ids := make([]int64, 0, len(recordHashes))
		frequencies := make(map[int64]int)
		indexes := make(map[int64]int64)
		if len(recordHashes) != 0 { // record hashes cannot be 0
			var id int64
			var val string
			rows, err := tx.Query(fmt.Sprintf(op.idQuery, strings.Join(recordHashes, "','")))
			if err != nil {
				return err
			}
			defer rows.Close()
			hash2ids := make(map[hash.Hash256]int64)
			for rows.Next() {
				if err := rows.Scan(&id, &val); err != nil {
					return err
				}
				h, err := hash.HexStringToHash256(val)
				if err != nil {
					return err
				}
				if h == hash.ZeroHash256 {
					continue
				}
				ids = append(ids, id)
				if hash2Frequencies[h] > 1 {
					frequencies[id] = hash2Frequencies[h]
				}
				hash2ids[h] = id
			}
			if err := rows.Err(); err != nil {
				return err
			}
			for i, h := range index2Hash {
				if h == hash.ZeroHash256 {
					continue
				}
				if id, ok := hash2ids[h]; ok {
					indexes[i] = id
				} else {
					return errors.Errorf("cannot find id for hash %x", h)
				}
			}
			if len(ids) != len(recordHashes) {
				return errors.New("wrong number of records")
			}
		}
		bidBytes, err := json.Marshal(ids)
		if err != nil {
			return err
		}
		freqBytes, err := json.Marshal(frequencies)
		if err != nil {
			return err
		}
		idxBytes, err := json.Marshal(indexes)
		if err != nil {
			return err
		}
		if _, err := tx.Exec(op.insertHeightToRecordsQuery, height, bidBytes, freqBytes, idxBytes); err != nil {
			return err
		}
	}

	return nil
}

// InsertDeltaBuckets inserts bucket records into table by tx
func InsertDeltaBuckets(tableName string, driverName DRIVERTYPE, records interface{}, tx *sql.Tx) (
	frequencies map[hash.Hash256]int,
	indexes map[int64]hash.Hash256,
	err error,
) {
	buckets, ok := records.([]*pyggBucket)
	if !ok {
		return nil, nil, errors.Errorf("invalid record type %s, *types.Bucket expected", reflect.TypeOf(records))
	}
	if buckets == nil {
		return nil, nil, nil
	}

	var stmt *sql.Stmt
	switch driverName {
	case SQLITE:
		stmt, err = tx.Prepare(fmt.Sprintf(InsertBucketsQuerySQLITE, tableName))
	case MYSQL:
		stmt, err = tx.Prepare(fmt.Sprintf(InsertBucketsQueryMySQL, tableName))
	default:
		return nil, nil, errors.New("wrong driver type")
	}
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		closeErr := stmt.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	frequencies = make(map[hash.Hash256]int)
	indexes = make(map[int64]hash.Hash256)
	for _, bucket := range buckets {
		h, err := bucket.Hash()
		if err != nil {
			return nil, nil, err
		}
		indexes[int64(bucket.Index)] = h
		if h == hash.ZeroHash256 {
			continue
		}
		if f, ok := frequencies[h]; ok {
			frequencies[h] = f + 1
		} else {
			frequencies[h] = 1
		}
		if _, err = stmt.Exec(
			hex.EncodeToString(h[:]),
			bucket.StartTime,
			bucket.Duration.String(),
			bucket.Amount.Bytes(),
			bucket.Decay,
			bucket.Owner.Bytes(),
			bucket.CanName[:],
		); err != nil {
			return nil, nil, err
		}
	}

	return frequencies, indexes, nil
}

// NewDeltaBucketTableOperator creates an operator for bucket table
func NewDeltaBucketTableOperator(tableName string, driverName DRIVERTYPE) (Operator, error) {
	var creation string
	switch driverName {
	case SQLITE:
		creation = "CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT, hash TEXT UNIQUE, start_time TIMESTAMP, duration TEXT, amount BLOB, decay INTEGER, voter BLOB, candidate BLOB)"
	case MYSQL:
		creation = "CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTO_INCREMENT, hash VARCHAR(64) UNIQUE, start_time TIMESTAMP, duration TEXT, amount BLOB, decay INTEGER, voter BLOB, candidate BLOB)"
	default:
		return nil, errors.New("Wrong driver type")
	}
	return NewDeltaRecordTableOperator(
		tableName,
		driverName,
		InsertDeltaBuckets,
		QueryBuckets,
		creation,
	)
}
