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
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type tableDB struct {
	db 			*sql.DB 
	path 		string
}


// NewTableStore creates a new tablestore
func NewTableDB(cfg Config) *tableDB {
	return &tableDB{
		path:       cfg.DBPath,
	}
}

// Start starts the tableDB 
func (t *tableDB) Start() error {
 	db, err := sql.Open("sqlite3", t.path)
 	if err != nil {
		return err
	}
	t.db = db
	return t.InitializeTable()
}

// InitializeTable initializes the tableDB 
func (t *tableDB) InitializeTable () error {
	if err := t.NewBucketTable(); err != nil {
		return err 
	}
	if err := t.NewRegistrationTable(); err != nil {
		return err
	}
	if err := t.NewHeightToRegTable(); err != nil {
		return err 
	}
	if err := t.NewHeightToBucketTable(); err != nil {
		return err
	}
	if err := t.NewHeightToTimeTable(); err != nil {
		return err
	}
	return	nil 
}

//NewBucketTable creates buckets table
func (t *tableDB) NewBucketTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS buckets (id INTEGER PRIMARY KEY AUTOINCREMENT, hash BLOB, startTime BLOB, duration TEXT, amount BLOB, decay INTEGER, voter BLOB, candidate BLOB, bucketIndex BLOB)")
	return err
}
//NewRegistrationTable creates registrations table 
func (t *tableDB) NewRegistrationTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS registrations (id INTEGER PRIMARY KEY AUTOINCREMENT, hash BLOB, name BLOB, address BLOB, operatorAddress BLOB, rewardAddress BLOB, selfStakingWeight TEXT)")
	return err
}
//NewHeightToRegTable creates heightToReg table
func (t *tableDB) NewHeightToRegTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS heightToReg (height BLOB, index INTEGER)")
	return err
}
//NewHeightToBucketTable creates heightToBucket table
func (t *tableDB) NewHeightToBucketTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS heightToBucket (height BLOB, index INTEGER)")
	return err
}
//NewHeightToTimeTable creates heightToTime table
func (t *tableDB) NewHeightToTimeTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS heightToTime (height BLOB, time BLOB)")
	return err
}

//Stop closes the tableDB
func (t *tableDB) Stop() error {
	return t.db.Close()
}
