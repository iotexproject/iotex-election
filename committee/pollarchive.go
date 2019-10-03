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
	"os"
	"reflect"
	"time"

	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"

	// require sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/types"
	"github.com/iotexproject/iotex-election/util"
)

// PollArchive stores registrations, buckets, and other data
type PollArchive interface {
	HeightBefore(time.Time) (uint64, error)
	// Buckets returns a list of Bucket of a given height
	Buckets(uint64) ([]*types.Bucket, error)
	// NativeBuckets returns a list of Bucket of a given epoch number
	NativeBuckets(uint64) ([]*types.Bucket, error)
	// Registrations returns a list of Registration of a given height
	Registrations(uint64) ([]*types.Registration, error)
	// MintTime returns the mint time of a given height
	MintTime(uint64) (time.Time, error)
	// NativeMintTime returns the mint time of a given epoch number
	NativeMintTime(uint64) (time.Time, error)
	// PutPoll puts one poll record
	PutPoll(uint64, time.Time, []*types.Registration, []*types.Bucket) error
	// PutNativePoll puts one native poll record on IoTeX chain
	PutNativePoll(uint64, time.Time, []*types.Bucket) error
	// TipHeight returns the tip height stored in archive
	TipHeight() (uint64, error)
	// Start starts the archive
	Start(context.Context) error
	// Stop stops the archive
	Stop(context.Context) error
}

type archive struct {
	startHeight               uint64
	interval                  uint64
	db                        *sql.DB
	bucketTableOperator       Operator
	nativeBucketTableOperator Operator
	registrationTableOperator Operator
	timeTableOperator         *TimeTableOperator
	nativeTimeTableOperator   *TimeTableOperator
	oldDB                     db.KVStoreWithNamespace
}

// NewArchive creates a new archive of poll
func NewArchive(dbPath string, numOfRetries uint8, startHeight uint64, interval uint64) (PollArchive, error) {
	fileExists := func(path string) bool {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			return false
		}
		if err != nil {
			zap.L().Panic("unexpected error", zap.Error(err))
		}
		return true
	}
	isOldCommitteeDB := func(oldDbPath string) bool {
		if !fileExists(oldDbPath) {
			return false
		}
		db, err := bolt.Open(oldDbPath, 0666, nil)
		if err != nil {
			if err == bolt.ErrInvalid {
				return false
			}
			zap.L().Panic("unexpected error", zap.Error(err))
		}
		if err = db.Close(); err != nil {
			zap.L().Panic("unexpected error", zap.Error(err))
		}
		return true
	}
	oldDbPath := dbPath + ".bolt"
	var kvstore db.KVStoreWithNamespace
	if isOldCommitteeDB(dbPath) {
		if err := os.Rename(dbPath, oldDbPath); err != nil {
			return nil, err
		}
	}
	if fileExists(oldDbPath) {
		kvstore = db.NewBoltDB(oldDbPath, numOfRetries)
	}
	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	bucketTableOperator, err := NewBucketTableOperator("buckets", true)
	if err != nil {
		return nil, err
	}
	nativeBucketTableOperator, err := NewBucketTableOperator("native_buckets", true)
	if err != nil {
		return nil, err
	}
	registrationTableOperator, err := NewRegistrationTableOperator("registrations", true)
	if err != nil {
		return nil, err
	}
	return &archive{
		db:                        sqlDB,
		startHeight:               startHeight,
		interval:                  interval,
		bucketTableOperator:       bucketTableOperator,
		nativeBucketTableOperator: nativeBucketTableOperator,
		registrationTableOperator: registrationTableOperator,
		timeTableOperator:         NewTimeTableOperator("mint_time", true),
		nativeTimeTableOperator:   NewTimeTableOperator("native_mint_time", true),
		oldDB:                     kvstore,
	}, nil
}

func (arch *archive) Registrations(height uint64) ([]*types.Registration, error) {
	value, err := arch.registrationTableOperator.Get(height, arch.db, nil)
	if err != nil {
		return nil, err
	}
	records, ok := value.([]*types.Registration)
	if !ok {
		return nil, errors.Errorf("Unexpected type %s", reflect.TypeOf(value))
	}
	return records, nil
}

func (arch *archive) Buckets(height uint64) ([]*types.Bucket, error) {
	value, err := arch.bucketTableOperator.Get(height, arch.db, nil)
	if err != nil {
		return nil, err
	}
	records, ok := value.([]*types.Bucket)
	if !ok {
		return nil, errors.Errorf("Unexpected type %s", reflect.TypeOf(value))
	}
	return records, nil
}

func (arch *archive) NativeBuckets(epochNum uint64) ([]*types.Bucket, error) {
	value, err := arch.nativeBucketTableOperator.Get(epochNum, arch.db, nil)
	if err != nil {
		return nil, err
	}
	records, ok := value.([]*types.Bucket)
	if !ok {
		return nil, errors.Errorf("Unexpected type %s", reflect.TypeOf(value))
	}
	return records, nil
}

func (arch *archive) PutPoll(height uint64, mintTime time.Time, regs []*types.Registration, buckets []*types.Bucket) (err error) {
	tx, err := arch.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := arch.registrationTableOperator.Put(height, regs, tx); err != nil {
		return err
	}
	if err := arch.bucketTableOperator.Put(height, buckets, tx); err != nil {
		return err
	}
	if err := arch.timeTableOperator.Put(height, mintTime, tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (arch *archive) PutNativePoll(epochNum uint64, mintTime time.Time, buckets []*types.Bucket) (err error) {
	tx, err := arch.db.Begin()
	if err != nil {
		return err
	}
	if err := arch.nativeBucketTableOperator.Put(epochNum, buckets, tx); err != nil {
		return err
	}
	if err := arch.nativeTimeTableOperator.Put(epochNum, mintTime, tx); err != nil {
		return err
	}
	defer tx.Rollback()
	return tx.Commit()
}

func (arch *archive) TipHeight() (uint64, error) {
	return arch.timeTableOperator.TipHeight(arch.db, nil)
}

func (arch *archive) HeightBefore(ts time.Time) (uint64, error) {
	return arch.timeTableOperator.HeightBefore(ts, arch.db, nil)
}

func (arch *archive) MintTime(height uint64) (time.Time, error) {
	value, err := arch.timeTableOperator.Get(height, arch.db, nil)
	if err != nil {
		return time.Time{}, err
	}
	mintTime, ok := value.(time.Time)
	if !ok {
		return time.Time{}, errors.Errorf("Unexpected type %s", reflect.TypeOf(value))
	}
	return mintTime, nil
}

func (arch *archive) NativeMintTime(epochNum uint64) (time.Time, error) {
	value, err := arch.nativeTimeTableOperator.Get(epochNum, arch.db, nil)
	if err != nil {
		return time.Time{}, err
	}
	mintTime, ok := value.(time.Time)
	if !ok {
		return time.Time{}, errors.Errorf("Unexpected type %s", reflect.TypeOf(value))
	}
	return mintTime, nil
}

func (arch *archive) Start(ctx context.Context) (err error) {
	var tx *sql.Tx
	tx, err = arch.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err = arch.bucketTableOperator.CreateTables(tx); err != nil {
		return err
	}
	if err = arch.nativeBucketTableOperator.CreateTables(tx); err != nil {
		return err
	}
	if err = arch.registrationTableOperator.CreateTables(tx); err != nil {
		return err
	}
	if err = arch.timeTableOperator.CreateTables(tx); err != nil {
		return err
	}
	if err = arch.nativeTimeTableOperator.CreateTables(tx); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return arch.migrate(ctx)
}

func (arch *archive) Stop(_ context.Context) (err error) {
	return nil
}

func (arch *archive) migrateResult(height uint64, r *types.ElectionResult) error {
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

	return arch.PutPoll(height, r.MintTime(), regs, buckets)
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
	tipHeight, err := arch.TipHeight()
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
