package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DBFactory struct {
	DriverName string

	DatabaseName string
}

func (db *DBFactory) GetConnection() *sql.DB {
	open, err := sql.Open(db.DriverName, db.DatabaseName)
	if err != nil {
		errors.New(fmt.Sprintf("database %s connect error : %s", db.DatabaseName, err.Error()))
	}
	return open
}
