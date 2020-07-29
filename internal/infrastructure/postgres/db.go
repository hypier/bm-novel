package postgres

import (
	"bm-novel/internal/config"
	"fmt"

	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"
)

// DefaultDB 数据库连接
var DefaultDB *sqlx.DB

func init() {
	config.LoadConfig()
	db, err := connectMysql()

	if err != nil {
		fmt.Println(err)
	}

	DefaultDB = db
}

func connectMysql() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Config.DB.UserName,
		config.Config.DB.Password, config.Config.DB.IPAddress, config.Config.DB.Port, config.Config.DB.DBName)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to user db")
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	err = db.Ping()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to connect to user db")
	}
	return db, nil
}
