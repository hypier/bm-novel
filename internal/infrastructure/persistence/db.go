package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joyparty/entity"
	_ "github.com/joyparty/entity"
)

var (
	userName  string = "root"
	password  string = "root"
	ipAddress string = "localhost"
	port      int    = 3306
	dbName    string = "db_admin"
	charset   string = "utf8"
)

func connectMysql() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	_ = db.Ping()

	return db, nil

}

func DoQuery(ctx context.Context, strSql string, ent entity.Entity, db *sqlx.DB) error {

	rows, err := sqlx.NamedQueryContext(ctx, db, strSql, ent)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return sql.ErrNoRows
	}

	if err := rows.StructScan(ent); err != nil {
		return fmt.Errorf("scan struct, %w", err)
	}

	return rows.Err()
}
