// Copyright (c) 2019 IoTeX
// This program is free software: you can redistribute it and/or modify it under the terms of the
// GNU General Public License as published by the Free Software Foundation, either version 3 of
// the License, or (at your option) any later version.
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY;
// without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See
// the GNU General Public License for more details.
// You should have received a copy of the GNU General Public License along with this program. If
// not, see <http://www.gnu.org/licenses/>.

package db

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"go.etcd.io/bbolt"
)

const (
	filemode = 0600
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
	NumOfRetries uint8  `yaml:"numOfRetries"`
	DBPath       string `yaml:"dbPath"`
}

// KVStore defines the db interface using in committee
type KVStore interface {
	Start(context.Context) error
	Stop(context.Context) error
	Get([]byte) ([]byte, error)
	Put([]byte, []byte) error
}

type memStore struct {
	kv map[string][]byte
}

// NewInMemKVStore creates a new in memory kv store
func NewInMemKVStore() KVStore {
	return &memStore{
		kv: make(map[string][]byte),
	}
}

// Start starts the in-memory store
func (m *memStore) Start(_ context.Context) error {
	return nil
}

// Stop stops hte in-memory store
func (m *memStore) Stop(_ context.Context) error {
	m.kv = nil
	return nil
}

// Get gets value by key from in-memory store
func (m *memStore) Get(key []byte) ([]byte, error) {
	value, ok := m.kv[string(key)]
	if !ok {
		return nil, errors.Wrapf(ErrNotExist, "key = %s", string(key))
	}
	return value, nil
}

// Put stores key and value to in-memory store
func (m *memStore) Put(key []byte, value []byte) error {
	m.kv[string(key)] = value
	return nil
}

// KVStoreWithNamespace defines the db interface with namesapce
type KVStoreWithNamespace interface {
	Start(context.Context) error
	Stop(context.Context) error
	Get(string, []byte) ([]byte, error)
	Put(string, []byte, []byte) error
}

// KVStoreWithNamespaceWrapper defines a wrapper to convert KVStoreWithNamespace to KVStore
type KVStoreWithNamespaceWrapper struct {
	store     KVStoreWithNamespace
	namespace string
}

// NewKVStoreWithNamespaceWrapper create a kvstore with specified namespace and a KVStoreWithNamespace
func NewKVStoreWithNamespaceWrapper(namespace string, store KVStoreWithNamespace) KVStore {
	return &KVStoreWithNamespaceWrapper{
		namespace: namespace,
		store:     store,
	}
}

// Start starts the kv store
func (w *KVStoreWithNamespaceWrapper) Start(ctx context.Context) error {
	return w.store.Start(ctx)
}

// Stop stops the kv store
func (w *KVStoreWithNamespaceWrapper) Stop(ctx context.Context) error {
	return w.store.Stop(ctx)
}

// Get gets the value by key from kv store
func (w *KVStoreWithNamespaceWrapper) Get(key []byte) ([]byte, error) {
	return w.store.Get(w.namespace, key)
}

// Put puts key-value pair into kv store
func (w *KVStoreWithNamespaceWrapper) Put(key []byte, value []byte) error {
	return w.store.Put(w.namespace, key, value)
}

type memStoreWithNamespace struct {
	kvs map[string]KVStore
}

// NewInMemStoreWithNamespace defines a wrapper to convert memStore to KVStoreWithNamespace
func NewInMemStoreWithNamespace(_ Config) KVStoreWithNamespace {
	return &memStoreWithNamespace{
		kvs: make(map[string]KVStore),
	}
}

// Start starts the memStoreWithNamespace
func (m *memStoreWithNamespace) Start(_ context.Context) error {
	return nil
}

// Stop stops the memStoreWithNamespace
func (m *memStoreWithNamespace) Stop(_ context.Context) error {
	m.kvs = nil
	return nil

}

// Get gets value by key from memStoreWithNamespace
func (m *memStoreWithNamespace) Get(namespace string, key []byte) ([]byte, error) {
	if kvstore, ok := m.kvs[namespace]; ok {
		return kvstore.Get(key)
	}
	//if the namespace doesn't exist, return error
	return nil, errors.New("Namespace doesn't exist")
}

// Put puts key and value to memStoreWithNamespace
func (m *memStoreWithNamespace) Put(namespace string, key []byte, value []byte) error {
	if kvstore, ok := m.kvs[namespace]; ok {
		return kvstore.Put(key, value)
	}
	//if the namespace doesn't exist, it makes new kvstore and namespace, and put the data
	newkvStore := NewInMemKVStore()
	if err := newkvStore.Put(key, value); err != nil {
		return err
	}
	m.kvs[namespace] = newkvStore
	return nil
}

type boltDB struct {
	db         *bbolt.DB
	path       string
	numRetries uint8
}

// NewBoltDB creates a new boltDB
func NewBoltDB(cfg Config) KVStoreWithNamespace {
	return &boltDB{
		numRetries: cfg.NumOfRetries,
		path:       cfg.DBPath,
	}
}

// Start starts the boltDB
func (b *boltDB) Start(_ context.Context) error {
	db, err := bbolt.Open(b.path, filemode, nil)
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
			fmt.Println(err)
			return errors.Wrap(ErrIO, err.Error())
		}
	}
	return nil
}

// Get gets value by key from boltDB
func (b *boltDB) Get(namespace string, key []byte) ([]byte, error) {
	var value []byte
	err := b.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(namespace))
		if bucket == nil {
			return errors.Wrapf(bbolt.ErrBucketNotFound, "bucket = %s", namespace)
		}
		value = bucket.Get(key)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if value == nil {
		err = errors.Wrapf(ErrNotExist, "key = %s", string(key))
	}
	return value, err
}

// Put stores key and value to boltDB
func (b *boltDB) Put(namespace string, key []byte, value []byte) error {
	var err error
	for c := uint8(0); c < b.numRetries; c++ {
		err = b.db.Update(func(tx *bbolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(namespace))
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
