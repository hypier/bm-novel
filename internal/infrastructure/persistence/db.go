package persistence

import (
	"bm-novel/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// DefaultDB 数据库连接
var DefaultDB *sqlx.DB

func init() {
	config.LoadConfig()
	DefaultDB, _ = connectMysql()
}

func connectMysql() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Config.DB.UserName,
		config.Config.DB.Password, config.Config.DB.IPAddress, config.Config.DB.Port, config.Config.DB.DBName)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	_ = db.Ping()

	return db, nil
}
