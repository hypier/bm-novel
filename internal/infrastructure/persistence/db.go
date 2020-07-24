package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joyparty/entity"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

var (
	userName  string = "postgres"
	password  string = "123456"
	ipAddress string = "localhost"
	port      int    = 5432
	dbName    string = "db_novel"
)

type rowMapper func(ent entity.Entity) entity.Entity

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

func DoQuery(ctx context.Context, strSql string, ent entity.Entity, db *sqlx.DB, mapper rowMapper) ([]entity.Entity, error) {

	var list []entity.Entity

	rows, err := sqlx.NamedQueryContext(ctx, db, strSql, ent)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fmt.Println(strSql)

	for rows.Next() {
		if err := rows.StructScan(ent); err != nil {
			return nil, fmt.Errorf("scan struct, %w", err)
		}

		list = append(list, mapper(ent))
	}

	return list, rows.Err()
}

func DoQueryOne(ctx context.Context, strSql string, ent entity.Entity, db *sqlx.DB) error {

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
