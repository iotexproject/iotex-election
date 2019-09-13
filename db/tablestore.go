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

type TableDB struct {
	db 			*sql.DB 
	path 		string
}


// NewTableStore creates a new tablestore
func NewTableDB(cfg Config) *TableDB {
	return &TableDB{
		path:       cfg.DBPath,
	}
}

// Start starts the tableDB 
func (t *TableDB) Start() error {
 	db, err := sql.Open("sqlite3", t.path)
 	if err != nil {
		return err
	}
	t.db = db
	return t.InitializeTable()
}

// InitializeTable initializes the tableDB 
func (t *TableDB) InitializeTable () error {
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
	if err := t.NewNextHeightTable(); err != nil {
		return err
	}
	if err := t.NewIdenticalBucketTable(); err != nil {
		return err
	}
	if err := t.NewIdenticalRegTable(); err != nil {
		return err
	}
	return	nil 
}

//NewBucketTable creates buckets table
func (t *TableDB) NewBucketTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS buckets (id INTEGER PRIMARY KEY AUTOINCREMENT, hash BLOB, startTime TIMESTAMP, duration TEXT, amount BLOB, decay INTEGER, voter BLOB, candidate BLOB, bucketIndex BLOB)")
	return err
}
//NewRegistrationTable creates registrations table 
func (t *TableDB) NewRegistrationTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS registrations (id INTEGER PRIMARY KEY AUTOINCREMENT, hash BLOB, name BLOB, address BLOB, operatorAddress BLOB, rewardAddress BLOB, selfStakingWeight INTEGER)")
	return err
}
//NewHeightToRegTable creates heightToReg table
func (t *TableDB) NewHeightToRegTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS heightToReg (height INTEGER, index INTEGER REFERENCES registrations(id), CONSTRAINT key PRIMARY KEY (height, index))")
	return err
}
//NewHeightToBucketTable creates heightToBucket table
func (t *TableDB) NewHeightToBucketTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS heightToBucket (id INTEGER PRIMARY KEY AUTOINCREMENT, height INTEGER, index INTEGER REFERENCES buckets(id), CONSTRAINT key PRIMARY KEY (height, index))")
	return err
}
//NewHeightToTimeTable creates heightToTime table
func (t *TableDB) NewHeightToTimeTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS heightToTime (height INTEGER PRIMARY KEY, time TIMESTAMP)")
	return err
}
//NewNextHeightTable creates nextHeight table
func (t *TableDB) NewNextHeightTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS nextHeight (key INTEGER PRIMARY KEY, height integer)")
	return err
}
//NewIdenticalBucketTable creates identicalbucket table 
func (t *TableDB) NewIdenticalBucketTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS identicalBucket (height INTEGER PRIMARY KEY, identicalheight INTEGER)")
	return err
}
//NewIdenticalRegTable creates identicalReg table
func (t *TableDB) NewIdenticalRegTable() error {
	_, err := t.db.Exec("CREATE TABLE IF NOT EXISTS identicalReg (height INTEGER PRIMARY KEY, identicalheight INTEGER)")
	return err
}

//Stop closes the tableDB
func (t *TableDB) Stop() error {
	return t.db.Close()
}
