// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package db

import (
	"context"

	"github.com/pkg/errors"
	"github.com/boltdb/bolt"
)

const (
	filemode = 0600
	electionNS = "elections"
)

var (
	// ErrIO indicates the generic error of DB I/O operation
	ErrIO = errors.New("DB I/O operation error")
	// ErrNotExist defines an error that the query has no return value in db
	ErrNotExist = errors.New("key does not exist")
	// NextHeightKey defines the constant key of next height
	NextHeightKey = []byte("next-height")
)

// Config defines the config of db
type Config struct {
	NumOfRetries              uint8  `yaml:"numOfRetries"`
	DBPath                    string `yaml:"dbPath"`
}

// KVStore defines the db interface using in committee
type KVStore interface {
	Start(context.Context) error
	Stop(context.Context) error
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
}

// NewKVStore creates a new key-value store
func NewKVStore(cfg Config) KVStore {
	if cfg.DBPath == "" || cfg.NumOfRetries < 1 {
		return &memStore{}
	}
	return &boltDB{numRetries: cfg.NumOfRetries, path: cfg.DBPath}
}

type memStore struct {
	kv map[string][]byte
}

// Start starts the in-memory store
func (m *memStore) Start(_ context.Context) error {
	m.kv = make(map[string][]byte)
	return nil
}

// Stop stops hte in-memory store
func (m *memStore) Stop(_ context.Context) error {
	m.kv = nil
	return nil
}

// Get gets value by key from in-memory store
func (s *memStore) Get(key []byte) ([]byte, error) {
	value, ok := s.kv[string(key)]
	if !ok {
		return nil, errors.Wrapf(ErrNotExist, "key = %s", string(key))
	}
	return value, nil
}

// Put stores key and value to in-memory store
func (s *memStore) Put(key []byte, value []byte) error {
	s.kv[string(key)] = value
	return nil
}

type boltDB struct {
	db *bolt.DB
	path string
	numRetries uint8
}

// Start starts the boltDB
func (b *boltDB) Start(_ context.Context) error {
	db, err := bolt.Open(b.path, filemode, nil)
	if err != nil {
		return errors.Wrapf(ErrIO, err.Error())
	}
	b.db = db
	return nil
}

// Stop stops the boltDB
func (b *boltDB) Stop(_ context.Context) error {
	if b.db != nil {
		if err := b.db.Close(); err != nil {
			return errors.Wrap(ErrIO, err.Error())
		}
	}
	return nil
}

// Get gets value by key from boltDB
func (b *boltDB) Get(key []byte) ([]byte, error) {
	var value []byte
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(electionNS))
		if bucket == nil {
			return errors.Wrapf(bolt.ErrBucketNotFound, "bucket = %s", electionNS)
		}
		value  = bucket.Get(key)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if value == nil {
		err = errors.Wrapf(ErrNotExist, "key = %s", string(key))
	}
	return value, nil
}

// Put stores key and value to boltDB
func (b *boltDB) Put(key []byte, value []byte) error {
	var err error
	for c := uint8(0); c < b.numRetries; c++ {
		err = b.db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(electionNS))
			if err != nil {
				return err
			}
			return bucket.Put(key, value)
		})

		if err == nil {
			break
		}
	}
	return err
}

