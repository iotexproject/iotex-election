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
	"reflect"
	"sync"
	"time"

	// require sqlite3 driver
	"github.com/pkg/errors"
	_ "modernc.org/sqlite"

	"github.com/iotexproject/iotex-election/types"
)

type BucketArchive struct {
	startHeight         uint64
	interval            uint64
	db                  *sql.DB
	bucketTableOperator Operator
	timeTableOperator   Operator

	// Put (native) polls are synchronized to get rid of the risk of reading uncommitted changes from other tx on the
	// same connection.
	mutex sync.Mutex
}

// NewBucketArchive creates a new archive of bucket
func NewBucketArchive(dbPath string, numOfRetries uint8, startHeight uint64, interval uint64) (*BucketArchive, error) {
	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	bucketTableOperator, err := NewDeltaBucketTableOperator("buckets", SQLITE)
	if err != nil {
		return nil, err
	}
	return &BucketArchive{
		db:                  sqlDB,
		startHeight:         startHeight,
		interval:            interval,
		timeTableOperator:   NewTimeTableOperator("mint_time", SQLITE),
		bucketTableOperator: bucketTableOperator,
	}, nil
}

func (arch *BucketArchive) TipHeight() (uint64, error) {
	return arch.timeTableOperator.TipHeight(arch.db, nil)
}

func (arch *BucketArchive) MintTime(height uint64) (time.Time, error) {
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

func (arch *BucketArchive) Buckets(height uint64) ([]*types.Bucket, error) {
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

func (arch *BucketArchive) PutDelta(height uint64, ts time.Time, updatedBuckets []*pyggBucket) (err error) {
	arch.mutex.Lock()
	defer arch.mutex.Unlock()

	tx, err := arch.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := arch.timeTableOperator.Put(height, ts, tx); err != nil {
		return err
	}
	if err := arch.bucketTableOperator.Put(height, updatedBuckets, tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (arch *BucketArchive) Start(ctx context.Context) (err error) {
	var tx *sql.Tx
	tx, err = arch.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err = arch.timeTableOperator.CreateTables(tx); err != nil {
		return err
	}
	if err = arch.bucketTableOperator.CreateTables(tx); err != nil {
		return err
	}
	return tx.Commit()
}

func (arch *BucketArchive) Stop(_ context.Context) (err error) {
	return arch.db.Close()
}
