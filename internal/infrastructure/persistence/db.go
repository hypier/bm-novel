package persistence

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	userName  = "postgres"
	password  = "123456"
	ipAddress = "localhost"
	port      = 5432
	dbName    = "db_novel"
)

func connectMysql() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", userName, password, ipAddress, port, dbName)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	_ = db.Ping()

	return db, nil
}
