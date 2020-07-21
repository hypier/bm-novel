package persistence

import (
	"bm-novel/internal/domain/user"
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql" // 载入 goqu mysql驱动
	"github.com/jmoiron/sqlx"
	"github.com/joyparty/entity"
	"github.com/pkg/errors"
)

var defaultDB *sqlx.DB

func init() {
	defaultDB, _ = connectMysql()
}

type UserRepository struct {
	ctx context.Context
}

func (u UserRepository) FindOne(id string) (*user.User, error) {
	usr := &user.User{UserId: id}
	if err := entity.Load(u.ctx, usr, defaultDB); err != nil {
		return nil, errors.New(err.Error())
	}
	return usr, nil
}

func (u UserRepository) FindByName(name string) (*user.User, error) {
	usr := &user.User{UserName: name}
	sql, params, err := goqu.From(usr.TableName()).Where(goqu.Ex{"user_name": name}).ToSQL()

	fmt.Println(sql, params, err)

	if err := DoQuery(u.ctx, sql, usr, defaultDB); err != nil {
		return nil, errors.New(err.Error())
	}
	return usr, nil
}

func (u UserRepository) Create(user *user.User) error {
	if _, err := entity.Insert(u.ctx, user, defaultDB); err != nil {
		return errors.New(err.Error())
	} else {
		return nil
	}
}

func (u UserRepository) Update(user *user.User) error {
	if err := entity.Update(u.ctx, user, defaultDB); err != nil {
		return errors.New(err.Error())
	} else {
		return nil
	}
}
