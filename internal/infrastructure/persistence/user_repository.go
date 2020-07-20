package persistence

import (
	"bm-novel/internal/domain/user"
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/joyparty/entity"
)

var defaultDB *sqlx.DB

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
	panic("implement me")
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
