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
	"math/big"
	"reflect"
	"strings"
	"time"

	// require sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/util"
)

// Operator defines an interface of operations on some tables in SQL DB
type Operator interface {
	// CreateTables prepares the tables for the operator
	CreateTables(*sql.Tx) error
	// Get returns the value by height
	Get(uint64, *sql.DB, *sql.Tx) (interface{}, error)
	// Put writes value for height
	Put(uint64, interface{}, *sql.Tx) error
}

// Record defines a record
type Record interface {
	Hash() (hash.Hash256, error)
}

// InsertRecordsFunc defines an api to insert records
type InsertRecordsFunc func(string, interface{}, *sql.Tx) (map[hash.Hash256]int, error)

// QueryRecordsFunc defines an api to query records
type QueryRecordsFunc func(string, map[int64]int, *sql.DB, *sql.Tx) (interface{}, error)

type recordTableOperator struct {
	tableName string

	frequencyQuery             string
	hashQuery                  string
	idQuery                    string
	identicalQuery             string
	insertHeightToRecordsQuery string
	insertIdenticalQuery       string
	tableCreations             []string

	insertRecordsFunc InsertRecordsFunc
	queryRecordsFunc  QueryRecordsFunc

	interval uint64
}

// NewBucketTableOperator creates an operator for bucket table
func NewBucketTableOperator(tableName string, interval uint64) (Operator, error) {
	return NewRecordTableOperator(
		tableName,
		interval,
		InsertBuckets,
		QueryBuckets,
		"CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT, hash TEXT UNIQUE, start_time TIMESTAMP, duration TEXT, amount BLOB, decay INTEGER, voter BLOB, candidate BLOB)",
	)
}

// NewRegistrationTableOperator create an operator for registration table
func NewRegistrationTableOperator(tableName string, interval uint64) (Operator, error) {
	return NewRecordTableOperator(
		tableName,
		interval,
		InsertRegistrations,
		QueryRegistrations,
		"CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT, hash TEXT UNIQUE, name BLOB, address BLOB, operator_address BLOB, reward_address BLOB, self_staking_weight INTEGER)",
	)
}

// NewRecordTableOperator creates a new arch of poll
func NewRecordTableOperator(
	tableName string,
	interval uint64,
	insertRecordsFunc InsertRecordsFunc,
	queryRecordsFunc QueryRecordsFunc,
	recordTableCreation string,
) (Operator, error) {
	return &recordTableOperator{
		tableName:                  tableName,
		interval:                   interval,
		frequencyQuery:             fmt.Sprintf("SELECT ids, frequencies FROM height_to_%s WHERE height = ?", tableName),
		hashQuery:                  fmt.Sprintf("SELECT id, hash FROM %s WHERE id IN (%s)", tableName, "%s"),
		idQuery:                    fmt.Sprintf("SELECT id, hash FROM %s WHERE hash IN ('%s')", tableName, "%s"),
		identicalQuery:             fmt.Sprintf("SELECT identical_to FROM identical_%s WHERE height = ?", tableName),
		insertHeightToRecordsQuery: fmt.Sprintf("INSERT OR REPLACE INTO height_to_%s (height, ids, frequencies) VALUES (?, ?, ?)", tableName),
		insertIdenticalQuery:       fmt.Sprintf("INSERT OR IGNORE INTO identical_%s (height, identical_to) VALUES (?, ?)", tableName),
		insertRecordsFunc:          insertRecordsFunc,
		queryRecordsFunc:           queryRecordsFunc,
		tableCreations: []string{
			fmt.Sprintf(recordTableCreation, tableName),
			fmt.Sprintf("CREATE TABLE IF NOT EXISTS height_to_%s (height INTEGER PRIMARY KEY, ids BLOB, frequencies BLOB)", tableName),
			fmt.Sprintf("CREATE TABLE IF NOT EXISTS identical_%s (height INTEGER PRIMARY KEY, identical_to INTEGER)", tableName),
		},
	}, nil
}

func (arch *recordTableOperator) Get(height uint64, db *sql.DB, tx *sql.Tx) (interface{}, error) {
	height, err := arch.identicalTo(height, db, tx)
	if err != nil {
		return nil, err
	}
	frequencies, err := arch.frequencies(height, db, tx)
	if err != nil {
		return nil, err
	}

	return arch.queryRecordsFunc(arch.tableName, frequencies, db, tx)
}

func (arch *recordTableOperator) Put(height uint64, records interface{}, tx *sql.Tx) (err error) {
	var hash2Frequencies map[hash.Hash256]int
	if hash2Frequencies, err = arch.insertRecordsFunc(arch.tableName, records, tx); err != nil {
		return err
	}
	lastIdenticalHeight, lastFrequencies, err := arch.hashes(height-arch.interval, nil, tx)
	if err != nil {
		return errors.Wrap(err, "failed to get record hashes")
	}
	if arch.hasIdenticalRecords(hash2Frequencies, lastFrequencies) {
		if _, err := tx.Exec(arch.insertIdenticalQuery, height, lastIdenticalHeight); err != nil {
			return err
		}
	} else {
		recordHashes := make([]string, 0, len(hash2Frequencies))
		for h := range hash2Frequencies {
			recordHashes = append(recordHashes, hex.EncodeToString(h[:]))
		}
		if len(recordHashes) != 0 {
			ids := make([]int64, 0, len(recordHashes))
			frequencies := make(map[int64]int)
			var id int64
			var val string
			rows, err := tx.Query(fmt.Sprintf(arch.idQuery, strings.Join(recordHashes, "','")))
			if err != nil {
				return err
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&id, &val); err != nil {
					return err
				}
				h, err := hash.HexStringToHash256(val)
				if err != nil {
					return err
				}
				ids = append(ids, id)
				if hash2Frequencies[h] > 1 {
					frequencies[id] = hash2Frequencies[h]
				}
			}
			if err := rows.Err(); err != nil {
				return err
			}
			if len(ids) != len(recordHashes) {
				return errors.New("wrong number of records")
			}
			bidBytes, err := json.Marshal(ids)
			if err != nil {
				return err
			}
			freqBytes, err := json.Marshal(frequencies)
			if err != nil {
				return err
			}
			if _, err := tx.Exec(arch.insertHeightToRecordsQuery, height, bidBytes, freqBytes); err != nil {
				return err
			}
		}
	}

	return nil
}

func (arch *recordTableOperator) CreateTables(tx *sql.Tx) (err error) {
	for _, creation := range arch.tableCreations {
		if _, err = tx.Exec(creation); err != nil {
			return err
		}
	}
	return nil
}

func (arch *recordTableOperator) hashes(height uint64, sdb *sql.DB, tx *sql.Tx) (uint64, map[hash.Hash256]int, error) {
	height, err := arch.identicalTo(height, sdb, tx)
	if err != nil {
		return 0, nil, err
	}
	frequencies, err := arch.frequencies(height, sdb, tx)
	switch err {
	case db.ErrNotExist:
		return height, nil, nil
	case nil:
	default:
		return 0, nil, err
	}
	ids := make([]int64, 0, len(frequencies))
	for bid := range frequencies {
		ids = append(ids, bid)
	}
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.Query(fmt.Sprintf(arch.hashQuery, atos(ids)))
	} else {
		rows, err = sdb.Query(fmt.Sprintf(arch.hashQuery, atos(ids)))
	}
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	hashes := make(map[hash.Hash256]int, len(ids))
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
		hashes[h] = frequencies[id]
	}
	if rows.Err() != nil {
		return 0, nil, rows.Err()
	}
	return height, hashes, nil
}

func (arch *recordTableOperator) frequencies(height uint64, sdb *sql.DB, tx *sql.Tx) (map[int64]int, error) {
	var (
		bidBytes, timeBytes []byte
		ids                 []int64
		err                 error
		frequencies         map[int64]int
	)
	if tx != nil {
		err = tx.QueryRow(arch.frequencyQuery, util.Uint64ToInt64(height)).Scan(&bidBytes, &timeBytes)
	} else {
		err = sdb.QueryRow(arch.frequencyQuery, util.Uint64ToInt64(height)).Scan(&bidBytes, &timeBytes)
	}
	switch err {
	case sql.ErrNoRows:
		return nil, db.ErrNotExist
	case nil:
	default:
		return nil, err
	}
	if err = json.Unmarshal(bidBytes, &ids); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(timeBytes, &frequencies); err != nil {
		return nil, err
	}
	bid2Times := make(map[int64]int, len(ids))
	for _, bid := range ids {
		f, ok := frequencies[bid]
		if !ok {
			bid2Times[bid] = 1
		} else {
			bid2Times[bid] = f
		}
	}
	return bid2Times, nil
}

func (arch *recordTableOperator) identicalTo(height uint64, sdb *sql.DB, tx *sql.Tx) (uint64, error) {
	var val int64
	var err error
	if tx != nil {
		err = tx.QueryRow(arch.identicalQuery, util.Uint64ToInt64(height)).Scan(&val)
	} else {
		err = sdb.QueryRow(arch.identicalQuery, util.Uint64ToInt64(height)).Scan(&val)
	}
	switch err {
	case nil:
		return uint64(val), nil
	case sql.ErrNoRows:
		return height, nil
	default:
		return 0, err
	}
}

func (arch *recordTableOperator) hasIdenticalRecords(
	records map[hash.Hash256]int,
	lastRecords map[hash.Hash256]int,
) bool {
	// if last height doesn't exist
	if lastRecords == nil {
		return false
	}
	// nil stands for identical
	if records == nil {
		return true
	}
	if len(records) != len(lastRecords) {
		return false
	}
	for h, last := range lastRecords {
		f, ok := records[h]
		if !ok {
			return false
		}
		if last != f {
			return false
		}
	}
	return true
}

// BucketRecordQuery is query to return buckets by ids
const BucketRecordQuery = "SELECT id, start_time, duration, amount, decay, voter, candidate FROM %s WHERE id IN (%s)"

// QueryBuckets returns buckets by ids
func QueryBuckets(tableName string, frequencies map[int64]int, sdb *sql.DB, tx *sql.Tx) (interface{}, error) {
	var (
		id, decay                int64
		startTime                time.Time
		rawDuration              string
		amount, voter, candidate []byte
		rows                     *sql.Rows
		err                      error
	)
	size := 0
	ids := make([]int64, 0, len(frequencies))
	for id, f := range frequencies {
		ids = append(ids, id)
		size += f
	}
	fmt.Println(tableName, BucketRecordQuery)
	if tx != nil {
		rows, err = tx.Query(fmt.Sprintf(BucketRecordQuery, tableName, atos(ids)))
	} else {
		rows, err = sdb.Query(fmt.Sprintf(BucketRecordQuery, tableName, atos(ids)))
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	buckets := make([]*types.Bucket, 0, size)
	for rows.Next() {
		if err := rows.Scan(&id, &startTime, &rawDuration, &amount, &decay, &voter, &candidate); err != nil {
			return nil, err
		}
		duration, err := time.ParseDuration(rawDuration)
		if err != nil {
			return nil, err
		}
		bucket, err := types.NewBucket(startTime, duration, big.NewInt(0).SetBytes(amount), voter, candidate, decay != 0)
		if err != nil {
			return nil, err
		}
		for i := frequencies[id]; i > 0; i-- {
			buckets = append(buckets, bucket)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return buckets, nil
}

// InsertBucketsQuery is query to insert buckets
const InsertBucketsQuery = "INSERT OR IGNORE INTO %s (hash, start_time, duration, amount, decay, voter, candidate) VALUES (?, ?, ?, ?, ?, ?, ?)"

// InsertBuckets inserts bucket records into table by tx
func InsertBuckets(tableName string, records interface{}, tx *sql.Tx) (frequencies map[hash.Hash256]int, err error) {
	var stmt *sql.Stmt
	if stmt, err = tx.Prepare(fmt.Sprintf(InsertBucketsQuery, tableName)); err != nil {
		return nil, err
	}
	defer func() {
		closeErr := stmt.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	frequencies = make(map[hash.Hash256]int)
	buckets, ok := records.([]*types.Bucket)
	if !ok {
		return nil, errors.Errorf("invalid record type %s, *types.Bucket expected", reflect.TypeOf(records))
	}
	for _, bucket := range buckets {
		h, err := bucket.Hash()
		if err != nil {
			return nil, err
		}
		if f, ok := frequencies[h]; ok {
			frequencies[h] = f + 1
		} else {
			frequencies[h] = 1
		}
		if _, err = stmt.Exec(
			hex.EncodeToString(h[:]),
			bucket.StartTime(),
			bucket.Duration().String(),
			bucket.Amount().Bytes(),
			bucket.Decay(),
			bucket.Voter(),
			bucket.Candidate(),
		); err != nil {
			return nil, err
		}
	}

	return frequencies, nil
}

// RegistrationQuery is query to get registrations by ids
const RegistrationQuery = "SELECT id, name, address, operator_address, reward_address, self_staking_weight FROM %s WHERE id IN (%s)"

// QueryRegistrations get all registrations by ids
func QueryRegistrations(tableName string, frequencies map[int64]int, sdb *sql.DB, tx *sql.Tx) (interface{}, error) {
	var (
		rows                                          *sql.Rows
		err                                           error
		name, address, operatorAddress, rewardAddress []byte
		id, selfStakingWeight                         int64
	)
	size := 0
	ids := make([]int64, 0, len(frequencies))
	for id, f := range frequencies {
		ids = append(ids, id)
		size += f
	}
	if tx != nil {
		rows, err = tx.Query(fmt.Sprintf(RegistrationQuery, tableName, atos(ids)))
	} else {
		rows, err = sdb.Query(fmt.Sprintf(RegistrationQuery, tableName, atos(ids)))
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	registrations := make([]*types.Registration, 0, size)
	for rows.Next() {
		if err := rows.Scan(&id, &name, &address, &operatorAddress, &rewardAddress, &selfStakingWeight); err != nil {
			return nil, err
		}
		registration := types.NewRegistration(name, address, operatorAddress, rewardAddress, uint64(selfStakingWeight))
		for i := frequencies[id]; i > 0; i-- {
			registrations = append(registrations, registration)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return registrations, nil
}

// InsertRegistrationQuery is query to insert registrations
const InsertRegistrationQuery = "INSERT OR IGNORE INTO %s (hash, name, address, operator_address, reward_address, self_staking_weight) VALUES (?, ?, ?, ?, ?, ?)"

// InsertRegistrations inserts registration records into table by tx
func InsertRegistrations(tableName string, records interface{}, tx *sql.Tx) (frequencies map[hash.Hash256]int, err error) {
	var regStmt *sql.Stmt
	if regStmt, err = tx.Prepare(fmt.Sprintf(InsertRegistrationQuery, tableName)); err != nil {
		return nil, err
	}
	defer func() {
		closeErr := regStmt.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	frequencies = make(map[hash.Hash256]int)
	regs, ok := records.([]*types.Registration)
	if !ok {
		return nil, errors.Errorf("Unexpected type %s", reflect.TypeOf(records))
	}
	for _, reg := range regs {
		var h hash.Hash256
		if h, err = reg.Hash(); err != nil {
			return nil, err
		}
		if f, ok := frequencies[h]; ok {
			frequencies[h] = f + 1
		} else {
			frequencies[h] = 1
		}
		if _, err = regStmt.Exec(
			hex.EncodeToString(h[:]),
			reg.Name(),
			reg.Address(),
			reg.OperatorAddress(),
			reg.RewardAddress(),
			util.Uint64ToInt64(reg.SelfStakingWeight()),
		); err != nil {
			return nil, err
		}
	}

	return frequencies, nil
}
