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
	"fmt"
	"reflect"
	"time"

	// require sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"

	"github.com/iotexproject/iotex-election/db"
	"github.com/iotexproject/iotex-election/util"
)

// TimeTableOperator defines an operator on timetable
type TimeTableOperator struct {
	createTableQuery    string
	heightQuery         string
	insertMintTimeQuery string
	mintTimeQuery       string
	tipHeightQuery      string
}

// NewTimeTableOperator returns an operator to time table
func NewTimeTableOperator(tableName string) *TimeTableOperator {
	return &TimeTableOperator{
		createTableQuery:    fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (height INTEGER PRIMARY KEY, time TIMESTAMP)", tableName),
		heightQuery:         fmt.Sprintf("SELECT MAX(height) FROM %s WHERE ? >= time AND EXISTS (SELECT * FROM %s WHERE ? <= time)", tableName, tableName),
		insertMintTimeQuery: fmt.Sprintf("INSERT OR IGNORE INTO %s (height, time) VALUES (?, ?)", tableName),
		mintTimeQuery:       fmt.Sprintf("SELECT time FROM %s WHERE height = ?", tableName),
		tipHeightQuery:      fmt.Sprintf("SELECT MAX(height) FROM %s", tableName),
	}
}

// TipHeight returns the tip height in the time table
func (operator *TimeTableOperator) TipHeight(sdb *sql.DB, tx *sql.Tx) (uint64, error) {
	var val sql.NullInt64
	var err error
	if tx != nil {
		err = tx.QueryRow(operator.tipHeightQuery).Scan(&val)
	} else {
		err = sdb.QueryRow(operator.tipHeightQuery).Scan(&val)
	}
	switch err {
	case sql.ErrNoRows:
		return 0, db.ErrNotExist
	case nil:
		if val.Valid {
			return uint64(val.Int64), nil
		}
		return 0, db.ErrNotExist
	default:
		return 0, err
	}
}

// HeightBefore returns the Height before ts in the time table
func (operator *TimeTableOperator) HeightBefore(ts time.Time, sdb *sql.DB, tx *sql.Tx) (height uint64, err error) {
	if tx != nil {
		err = tx.QueryRow(operator.heightQuery, ts, ts).Scan(&height)
	} else {
		err = sdb.QueryRow(operator.heightQuery, ts, ts).Scan(&height)
	}
	return uint64(height), nil
}

// CreateTables prepares the tables for the operator
func (operator *TimeTableOperator) CreateTables(tx *sql.Tx) (err error) {
	_, err = tx.Exec(operator.createTableQuery)

	return err
}

// Get returns the value by height
func (operator *TimeTableOperator) Get(height uint64, sdb *sql.DB, tx *sql.Tx) (interface{}, error) {
	var val time.Time
	var err error
	if tx != nil {
		err = tx.QueryRow(operator.mintTimeQuery, util.Uint64ToInt64(height)).Scan(&val)
	} else {
		err = sdb.QueryRow(operator.mintTimeQuery, util.Uint64ToInt64(height)).Scan(&val)
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

// Put writes value for height
func (operator *TimeTableOperator) Put(height uint64, value interface{}, tx *sql.Tx) error {
	mintTime, ok := value.(time.Time)
	if !ok {
		return errors.Errorf("unexpected type %s", reflect.TypeOf(value))
	}
	_, err := tx.Exec(operator.insertMintTimeQuery, height, mintTime)

	return err
}
