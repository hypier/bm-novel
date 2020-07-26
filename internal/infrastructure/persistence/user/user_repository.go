package user

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/persistence"
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/joyparty/entity"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type UserRepository struct {
	Ctx context.Context
}

func (u *UserRepository) FindList(roleCode []string, realName string, pageIndex int, pageSize int) (user.Users, error) {

	var r pq.StringArray = roleCode
	usr := &user.User{}

	var expressions []exp.Expression

	if roleCode != nil {
		expressions = append(expressions, goqu.L(`role_code @> ?`, r))
	}

	if realName != "" {
		expressions = append(expressions, goqu.L(`real_name = ?`, realName))
	}

	if pageSize == 0 {
		pageSize = 10
	}
	if pageIndex > 0 {
		pageIndex = pageIndex - 1
	}

	offset := pageSize * pageIndex

	strSQL, params, err := goqu.From(usr.TableName()).
		Where(expressions...).
		Limit(uint(pageSize)).
		Offset(uint(offset)).Order(goqu.I("create_at").Desc()).ToSQL()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	fmt.Println(strSQL)

	users := &user.Users{}
	err = persistence.DefaultDB.SelectContext(u.Ctx, users, strSQL, params...)

	return *users, err
}

func (u *UserRepository) FindOne(id string) (*user.User, error) {
	userID, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	usr := &user.User{UserID: userID}
	if err := entity.Load(u.Ctx, usr, persistence.DefaultDB); err != nil {
		return nil, errors.New(err.Error())
	}

	usr.SetPersistence()
	return usr, nil
}

func (u *UserRepository) FindByName(name string) (*user.User, error) {
	usr := user.User{}
	strSQL, params, err := goqu.From(usr.TableName()).Where(goqu.Ex{"user_name": name}).ToSQL()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	users := &user.Users{}
	err = persistence.DefaultDB.SelectContext(u.Ctx, users, strSQL, params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	for _, v := range *users {
		v.SetPersistence()
		return v, nil
	}

	return nil, err
}

func (u *UserRepository) Create(user *user.User) error {
	if _, err := entity.Insert(u.Ctx, user, persistence.DefaultDB); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (u *UserRepository) Update(user *user.User) error {
	if err := entity.Update(u.Ctx, user, persistence.DefaultDB); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
