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
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	"go.uber.org/zap"
	// require sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/util"
)

var tableCreations = []string{
	"CREATE TABLE IF NOT EXISTS buckets (id INTEGER PRIMARY KEY AUTOINCREMENT, hash TEXT UNIQUE, start_time TIMESTAMP, duration TEXT, amount BLOB, decay INTEGER, voter BLOB, candidate BLOB)",
	"CREATE TABLE IF NOT EXISTS registrations (id INTEGER PRIMARY KEY AUTOINCREMENT, hash TEXT UNIQUE, name BLOB, address BLOB, operator_address BLOB, reward_address BLOB, self_staking_weight INTEGER)",
	"CREATE TABLE IF NOT EXISTS height_to_registrations (height INTEGER, rid INTEGER REFERENCES registrations(id), CONSTRAINT key PRIMARY KEY (height, rid))",
	"CREATE TABLE IF NOT EXISTS height_to_buckets (height INTEGER PRIMARY KEY, bids BLOB, frequencies BLOB)",
	"CREATE TABLE IF NOT EXISTS mint_times (height INTEGER PRIMARY KEY, time TIMESTAMP)",
	"CREATE TABLE IF NOT EXISTS identical_buckets (height INTEGER PRIMARY KEY, identical_to INTEGER)",
	"CREATE TABLE IF NOT EXISTS identical_registrations (height INTEGER PRIMARY KEY, identical_to INTEGER)",
}

// Archive stores registrations, buckets, and other data
type Archive interface {
	HeightBefore(time.Time) (uint64, error)
	// Buckets returns a list of Bucket of a given height
	Buckets(uint64) ([]*types.Bucket, error)
	// Registrations returns a list of Registration of a given height
	Registrations(uint64) ([]*types.Registration, error)
	// MintTime returns the mint time of a given height
	MintTime(uint64) (time.Time, error)
	// PutPoll puts one poll record
	PutPoll(uint64, time.Time, []*types.Registration, []*types.Bucket) error
	// PutPolls puts multiple poll record
	PutPolls([]uint64, []time.Time, [][]*types.Registration, [][]*types.Bucket) error
	// TipHeight returns the tip height stored in archive
	TipHeight() (uint64, error)
	// Start starts the archive
	Start(context.Context) error
	// Stop stops the archive
	Stop(context.Context) error
}

type archive struct {
	startHeight uint64
	interval    uint64
	db          *sql.DB
	oldDB       db.KVStoreWithNamespace
}

// NewArchive creates a new arch of poll
func NewArchive(newDB *sql.DB, startHeight uint64, interval uint64, oldDB db.KVStoreWithNamespace) (Archive, error) {
	return &archive{
		db:          newDB,
		startHeight: startHeight,
		interval:    interval,
		oldDB:       oldDB,
	}, nil
}

func (arch *archive) Registrations(height uint64) ([]*types.Registration, error) {
	return arch.registrations(height, nil)
}

func (arch *archive) Buckets(height uint64) ([]*types.Bucket, error) {
	return arch.buckets(height, nil)
}

func (arch *archive) PutPoll(height uint64, mintTime time.Time, regs []*types.Registration, buckets []*types.Bucket) (err error) {
	var tx *sql.Tx
	if tx, err = arch.db.Begin(); err != nil {
		return err
	}
	defer tx.Rollback()
	if err = arch.put(height, mintTime, regs, buckets, tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (arch *archive) PutPolls(heights []uint64, mintTimes []time.Time, regs [][]*types.Registration, buckets [][]*types.Bucket) (err error) {
	indice := map[uint64]int{}
	for i, height := range heights {
		if _, ok := indice[height]; ok {
			return errors.Errorf("duplicate height %d", height)
		}
		indice[height] = i
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	var tx *sql.Tx
	if tx, err = arch.db.Begin(); err != nil {
		return err
	}
	defer tx.Rollback()
	for _, height := range heights {
		index := indice[height]
		if err = arch.put(height, mintTimes[index], regs[index], buckets[index], tx); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (arch *archive) TipHeight() (uint64, error) {
	return arch.tipHeight(nil)
}

const heightQuery = "SELECT MAX(height) FROM mint_times WHERE ? >= time AND EXISTS (SELECT * FROM mint_times WHERE ? <= time)"
const tipHeightQuery = "SELECT MAX(height) FROM mint_times"

func (arch *archive) HeightBefore(ts time.Time) (uint64, error) {
	var height int64
	if err := arch.db.QueryRow(heightQuery, ts, ts).Scan(&height); err != nil {
		return 0, err
	}
	return uint64(height), nil
}

func (arch *archive) MintTime(height uint64) (time.Time, error) {
	return arch.mintTime(height, nil)
}

func (arch *archive) Start(ctx context.Context) (err error) {
	if err = arch.createTables(); err != nil {
		return err
	}
	return arch.migrate(ctx)
}

func (arch *archive) Stop(_ context.Context) (err error) {
	return nil
}

func (arch *archive) createTables() (err error) {
	var tx *sql.Tx
	tx, err = arch.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, creation := range tableCreations {
		if _, err = tx.Exec(creation); err != nil {
			return err
		}
	}
	return tx.Commit()
}

func (arch *archive) migrateResult(height uint64, r *types.ElectionResult) error {
	tx, err := arch.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	candidates := r.Delegates()
	regs := make([]*types.Registration, 0, len(candidates))
	for _, candidate := range candidates {
		regs = append(regs, &candidate.Registration)
	}
	votes := r.Votes()
	buckets := make([]*types.Bucket, 0, len(votes))
	for _, vote := range votes {
		buckets = append(buckets, &vote.Bucket)
	}
	if err := arch.put(height, r.MintTime(), regs, buckets, tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (arch *archive) migrate(ctx context.Context) (err error) {
	if arch.oldDB == nil {
		return nil
	}
	kvstore := db.NewKVStoreWithNamespaceWrapper("electionNS", arch.oldDB)
	if err = kvstore.Start(ctx); err != nil {
		return err
	}
	defer func() {
		stopErr := kvstore.Stop(ctx)
		if err == nil {
			err = stopErr
		}
	}()
	nextHeightHash, err := kvstore.Get(db.NextHeightKey)
	if err != nil {
		return err
	}
	nextHeight := util.BytesToUint64(nextHeightHash)
	tipHeight, err := arch.tipHeight(nil)
	var heightToFetch uint64
	if err != nil {
		heightToFetch = arch.startHeight
	} else {
		heightToFetch = tipHeight + arch.interval
		if heightToFetch >= nextHeight {
			return nil
		}
	}
	lastPrintTime := time.Time{}
	for heightToFetch < nextHeight {
		if time.Now().Sub(lastPrintTime) > 5*time.Second {
			zap.L().Info("migrating", zap.Uint64("height", heightToFetch))
			lastPrintTime = time.Now()
		}
		data, err := kvstore.Get(util.Uint64ToBytes(heightToFetch))
		if err != nil {
			return err
		}
		r := &types.ElectionResult{}
		if err = r.Deserialize(data); err != nil {
			return err
		}
		if err = arch.migrateResult(heightToFetch, r); err != nil {
			return errors.Wrapf(err, "failed to migrate %d", heightToFetch)
		}
		heightToFetch += arch.interval
	}
	return nil
}

const (
	bucketFrequencyQuery       = "SELECT bids, frequencies FROM height_to_buckets WHERE height = ?"
	bucketHashQuery            = "SELECT id, hash FROM buckets WHERE id IN (%s)"
	bucketIDQuery              = "SELECT id, hash FROM buckets WHERE hash IN ('%s')"
	bucketQuery                = "SELECT id, start_time, duration, amount, decay, voter, candidate FROM buckets WHERE id IN (%s)"
	identicalBucketQuery       = "SELECT identical_to FROM identical_buckets WHERE height = ?"
	identicalRegistrationQuery = "SELECT identical_to FROM identical_registrations WHERE height = ?"
	insertBucketQuery          = "INSERT OR IGNORE INTO buckets (hash, start_time, duration, amount, decay, voter, candidate) VALUES (?, ?, ?, ?, ?, ?, ?)"
	insertIdenticalBucketQuery = "INSERT OR IGNORE INTO identical_buckets (height, identical_to) VALUES (?, ?)"
	insertIndenticalRegQuery   = "INSERT OR IGNORE INTO identical_registrations (height, identical_to) VALUES (?, ?)"
	insertMintTimeQuery        = "INSERT OR IGNORE INTO mint_times (height, time) VALUES (?, ?)"
	insertRegQuery             = "INSERT OR IGNORE INTO registrations (hash, name, address, operator_address, reward_address, self_staking_weight) VALUES (?, ?, ?, ?, ?, ?)"
	insertHeightToBucketsQuery = "INSERT OR REPLACE INTO height_to_buckets (height, bids, frequencies) VALUES (?, ?, ?)"
	mintTimeQuery              = "SELECT time FROM mint_times WHERE height = ?"
	registrationHashQuery      = "SELECT r.hash FROM registrations as r INNER JOIN height_to_registrations as hr ON r.id = hr.rid WHERE hr.height = ?"
	registrationQuery          = `SELECT r.name, r.address, r.operator_address, r.reward_address, r.self_staking_weight
								FROM registrations as r INNER JOIN height_to_registrations as hr
								ON r.id = hr.rid WHERE hr.height = ?`
)

func (arch *archive) registrationHashes(height uint64, tx *sql.Tx) (uint64, []hash.Hash256, error) {
	height, err := arch.heightWithIdenticalRegs(height, tx)
	if err != nil {
		return 0, nil, err
	}
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.Query(registrationHashQuery, util.Uint64ToInt64(height))
	} else {
		rows, err = arch.db.Query(registrationHashQuery, util.Uint64ToInt64(height))
	}
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	hashes := []hash.Hash256{}
	for rows.Next() {
		var val string
		if err := rows.Scan(&val); err != nil {
			return 0, nil, err
		}
		h, err := hash.HexStringToHash256(val)
		if err != nil {
			return 0, nil, err
		}
		hashes = append(hashes, h)
	}
	if rows.Err() != nil {
		return 0, nil, rows.Err()
	}
	return height, hashes, nil
}

func (arch *archive) mintTime(height uint64, tx *sql.Tx) (time.Time, error) {
	var val time.Time
	var err error
	if tx != nil {
		err = tx.QueryRow(mintTimeQuery, util.Uint64ToInt64(height)).Scan(&val)
	} else {
		err = arch.db.QueryRow(mintTimeQuery, util.Uint64ToInt64(height)).Scan(&val)
	}
	switch err {
	case sql.ErrNoRows:
		return time.Time{}, db.ErrNotExist
	case nil:
		return val, nil
	default:
		return time.Time{}, err
	}
}

func (arch *archive) tipHeight(tx *sql.Tx) (uint64, error) {
	var val int64
	var err error
	if tx != nil {
		err = tx.QueryRow(tipHeightQuery).Scan(&val)
	} else {
		err = arch.db.QueryRow(tipHeightQuery).Scan(&val)
	}
	switch err {
	case sql.ErrNoRows:
		return 0, db.ErrNotExist
	case nil:
		return uint64(val), nil
	default:
		return 0, err
	}
}

func (arch *archive) registrations(height uint64, tx *sql.Tx) ([]*types.Registration, error) {
	height, err := arch.heightWithIdenticalRegs(height, tx)
	if err != nil {
		return nil, err
	}
	var name, address, operatorAddress, rewardAddress []byte
	var selfStakingWeight int64
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.Query(registrationQuery, util.Uint64ToInt64(height))
	} else {
		rows, err = arch.db.Query(registrationQuery, util.Uint64ToInt64(height))
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	registrations := []*types.Registration{}
	for rows.Next() {
		if err := rows.Scan(&name, &address, &operatorAddress, &rewardAddress, &selfStakingWeight); err != nil {
			return nil, err
		}
		reg := types.NewRegistration(name, address, operatorAddress, rewardAddress, uint64(selfStakingWeight))
		registrations = append(registrations, reg)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return registrations, nil
}

func (arch *archive) buckets(height uint64, tx *sql.Tx) (buckets []*types.Bucket, err error) {
	height, err = arch.heightWithIdenticalBuckets(height, tx)
	if err != nil {
		return
	}
	var id, decay int64
	var startTime time.Time
	var rawDuration string
	var amount, voter, candidate []byte
	frequencies, err := arch.bucketFrequencies(height, tx)
	if err != nil {
		return nil, err
	}
	bids := make([]int64, 0, len(frequencies))
	for bid := range frequencies {
		bids = append(bids, bid)
	}
	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.Query(fmt.Sprintf(bucketQuery, atos(bids)))
	} else {
		rows, err = arch.db.Query(fmt.Sprintf(bucketQuery, atos(bids)))
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id, &startTime, &rawDuration, &amount, &decay, &voter, &candidate); err != nil {
			return nil, err
		}
		duration, err := time.ParseDuration(rawDuration)
		bucket, err := types.NewBucket(startTime, duration, big.NewInt(0).SetBytes(amount), voter, candidate, decay != 0)
		if err != nil {
			return nil, err
		}
		for i := frequencies[id]; i > 0; i-- {
			buckets = append(buckets, bucket)
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return buckets, nil
}

func (arch *archive) bucketHashes(height uint64, tx *sql.Tx) (uint64, map[hash.Hash256]int, error) {
	height, err := arch.heightWithIdenticalBuckets(height, tx)
	if err != nil {
		return 0, nil, err
	}
	frequencies, err := arch.bucketFrequencies(height, tx)
	switch err {
	case db.ErrNotExist:
		return height, nil, nil
	case nil:
	default:
		return 0, nil, err
	}
	bids := make([]int64, 0, len(frequencies))
	for bid := range frequencies {
		bids = append(bids, bid)
	}
	rows, err := arch.db.Query(fmt.Sprintf(bucketHashQuery, atos(bids)))
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	hashes := make(map[hash.Hash256]int, len(bids))
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

func (arch *archive) bucketFrequencies(height uint64, tx *sql.Tx) (map[int64]int, error) {
	var (
		bidBytes, timeBytes []byte
		bids                []int64
		err                 error
		frequencies         map[int64]int
	)
	if tx != nil {
		err = tx.QueryRow(bucketFrequencyQuery, util.Uint64ToInt64(height)).Scan(&bidBytes, &timeBytes)
	} else {
		err = arch.db.QueryRow(bucketFrequencyQuery, util.Uint64ToInt64(height)).Scan(&bidBytes, &timeBytes)
	}
	switch err {
	case sql.ErrNoRows:
		return nil, db.ErrNotExist
	case nil:
	default:
		return nil, err
	}
	if err = json.Unmarshal(bidBytes, &bids); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(timeBytes, &frequencies); err != nil {
		return nil, err
	}
	bid2Times := make(map[int64]int, len(bids))
	for _, bid := range bids {
		f, ok := frequencies[bid]
		if !ok {
			bid2Times[bid] = 1
		} else {
			bid2Times[bid] = f
		}
	}
	return bid2Times, nil
}

func (arch *archive) heightWithIdenticalRegs(height uint64, tx *sql.Tx) (uint64, error) {
	var val int64
	var err error
	if tx != nil {
		err = tx.QueryRow(identicalRegistrationQuery, util.Uint64ToInt64(height)).Scan(&val)
	} else {
		err = arch.db.QueryRow(identicalRegistrationQuery, util.Uint64ToInt64(height)).Scan(&val)
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

func (arch *archive) heightWithIdenticalBuckets(height uint64, tx *sql.Tx) (uint64, error) {
	var val int64
	var err error
	if tx != nil {
		err = tx.QueryRow(identicalBucketQuery, util.Uint64ToInt64(height)).Scan(&val)
	} else {
		err = arch.db.QueryRow(identicalBucketQuery, util.Uint64ToInt64(height)).Scan(&val)
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

func (arch *archive) hasIdenticalRegistrations(
	regs []*types.Registration,
	lastRegHashes []hash.Hash256,
) bool {
	// if last height doesn't exist
	if lastRegHashes == nil {
		return false
	}
	// nil stands for identical
	if regs == nil {
		return true
	}
	rhs := map[hash.Hash256]bool{}
	for _, reg := range regs {
		h, err := reg.Hash()
		if err != nil {
			return false
		}
		if _, ok := rhs[h]; ok {
			return false
		}
		rhs[h] = true
	}
	if len(rhs) != len(lastRegHashes) {
		return false
	}
	for _, h := range lastRegHashes {
		if _, ok := rhs[h]; !ok {
			return false
		}
		rhs[h] = false
	}
	return true
}

func (arch *archive) hasIdenticalBuckets(
	buckets map[hash.Hash256]int,
	lastBuckets map[hash.Hash256]int,
) bool {
	// if last height doesn't exist
	if lastBuckets == nil {
		return false
	}
	// nil stands for identical
	if buckets == nil {
		return true
	}
	if len(buckets) != len(lastBuckets) {
		return false
	}
	for h, last := range lastBuckets {
		f, ok := buckets[h]
		if !ok {
			return false
		}
		if last != f {
			return false
		}
	}
	return true
}

func (arch *archive) put(height uint64, mintTime time.Time, regs []*types.Registration, buckets []*types.Bucket, tx *sql.Tx) error {
	if err := arch.putRegistrations(height, regs, tx); err != nil {
		return errors.Wrap(err, "failed to store registrations")
	}
	if err := arch.putBuckets(height, buckets, tx); err != nil {
		return errors.Wrap(err, "failed to store buckets")
	}
	if err := arch.putMintTime(height, mintTime, tx); err != nil {
		return errors.Wrap(err, "failed to put mint time")
	}
	return nil
}

func (arch *archive) putRegistrations(height uint64, regs []*types.Registration, tx *sql.Tx) (err error) {
	var regStmt *sql.Stmt
	if regStmt, err = tx.Prepare(insertRegQuery); err != nil {
		return err
	}
	defer func() {
		closeErr := regStmt.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	for _, reg := range regs {
		var h hash.Hash256
		if h, err = reg.Hash(); err != nil {
			return err
		}
		if _, err = regStmt.Exec(
			hex.EncodeToString(h[:]),
			reg.Name(),
			reg.Address(),
			reg.OperatorAddress(),
			reg.RewardAddress(),
			util.Uint64ToInt64(reg.SelfStakingWeight()),
		); err != nil {
			return err
		}
	}
	irh, lastRegHashes, err := arch.registrationHashes(height-arch.interval, tx)
	if err != nil {
		return err
	}
	if arch.hasIdenticalRegistrations(regs, lastRegHashes) {
		if _, err = tx.Exec(insertIndenticalRegQuery, height, irh); err != nil {
			return err
		}
	} else {
		if _, err = tx.Exec("DROP TABLE IF EXISTS temp_regs"); err != nil {
			return err
		}
		if _, err = tx.Exec("CREATE TABLE temp_regs (height INTEGER, hash TEXT PRIMARY KEY)"); err != nil {
			return err
		}
		stmt, err := tx.Prepare("INSERT INTO temp_regs (height, hash) VALUES (?, ?)")
		if err != nil {
			return err
		}
		defer stmt.Close()
		for _, reg := range regs {
			h, err := reg.Hash()
			if err != nil {
				return err
			}
			if _, err = stmt.Exec(height, hex.EncodeToString(h[:])); err != nil {
				return err
			}
		}
		result, err := tx.Exec(`INSERT OR REPLACE INTO height_to_registrations (height, rid) 
			SELECT temp_regs.height, registrations.id FROM registrations INNER JOIN temp_regs ON registrations.hash=temp_regs.hash
		`)
		if err != nil {
			return err
		}
		rows, err := result.RowsAffected()
		if err != nil {
			return err
		}
		if rows != int64(len(regs)) {
			return errors.New("wrong number of registration records")
		}
		if _, err := tx.Exec("DROP TABLE temp_regs"); err != nil {
			return err
		}
	}
	return nil
}

func (arch *archive) putBuckets(height uint64, buckets []*types.Bucket, tx *sql.Tx) (err error) {
	var bucketStmt *sql.Stmt
	if bucketStmt, err = tx.Prepare(insertBucketQuery); err != nil {
		return err
	}
	defer func() {
		closeErr := bucketStmt.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	var h hash.Hash256
	for _, bucket := range buckets {
		if h, err = bucket.Hash(); err != nil {
			return err
		}
		if _, err = bucketStmt.Exec(
			hex.EncodeToString(h[:]),
			bucket.StartTime(),
			bucket.Duration().String(),
			bucket.Amount().Bytes(),
			bucket.Decay(),
			bucket.Voter(),
			bucket.Candidate(),
		); err != nil {
			return err
		}
	}

	ibh, lastbucketHashes, err := arch.bucketHashes(height-arch.interval, tx)
	if err != nil {
		return errors.Wrap(err, "failed to get bucket hashes")
	}
	var bh2Frequencies map[hash.Hash256]int
	var bucketHashes []string
	if buckets != nil {
		bh2Frequencies = make(map[hash.Hash256]int)
		bucketHashes = []string{}
		for _, bucket := range buckets {
			h, err := bucket.Hash()
			if err != nil {
				return err
			}
			if f, ok := bh2Frequencies[h]; ok {
				bh2Frequencies[h] = f + 1
			} else {
				bucketHashes = append(bucketHashes, hex.EncodeToString(h[:]))
				bh2Frequencies[h] = 1
			}
		}
	}
	if arch.hasIdenticalBuckets(bh2Frequencies, lastbucketHashes) {
		if _, err := tx.Exec(insertIdenticalBucketQuery, height, ibh); err != nil {
			return err
		}
	} else {
		if len(bucketHashes) != 0 {
			ids := make([]int64, 0, len(bucketHashes))
			frequencies := make(map[int64]int)
			var id int64
			var val string
			rows, err := tx.Query(fmt.Sprintf(bucketIDQuery, strings.Join(bucketHashes, "','")))
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
				if bh2Frequencies[h] > 1 {
					frequencies[id] = bh2Frequencies[h]
				}
			}
			if err := rows.Err(); err != nil {
				return err
			}
			if len(ids) != len(bucketHashes) {
				return errors.New("wrong number of bucket records")
			}
			bidBytes, err := json.Marshal(ids)
			if err != nil {
				return err
			}
			timeBytes, err := json.Marshal(frequencies)
			if err != nil {
				return err
			}
			if _, err := tx.Exec(insertHeightToBucketsQuery, height, bidBytes, timeBytes); err != nil {
				return err
			}
		}
	}

	return nil
}

func (arch *archive) putMintTime(height uint64, mintTime time.Time, tx *sql.Tx) error {
	if _, err := tx.Exec(insertMintTimeQuery, height, mintTime); err != nil {
		return err
	}
	return nil
}
