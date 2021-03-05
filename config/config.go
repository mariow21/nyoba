package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username string = "PharmanetBois"
	username1 string = "13.250.197.171"
	password string = "d3v3l0p8015"
	tcp string = "3306"
	database string = "test"
)

var (
	dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, username1, tcp, database)
)
// PharmanetBois:d3v3l0p8015@tcp(13.250.197.171:3306)/test

func MySQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	return db, nil
}
