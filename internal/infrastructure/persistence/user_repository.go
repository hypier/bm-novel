package persistence

import (
	"bm-novel/internal/domain/user"
	"context"
	"database/sql"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jmoiron/sqlx"
	"github.com/joyparty/entity"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var defaultDB *sqlx.DB

func init() {
	defaultDB, _ = connectMysql()
}

type UserRepository struct {
	Ctx context.Context
}

func (u *UserRepository) FindList(roleCode []string, realName string, pageIndex int, pageSize int) ([]user.User, error) {

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

	strSql, _, err := goqu.From(usr.TableName()).
		Where(expressions...).
		Limit(uint(pageSize)).
		Offset(uint(offset)).Order(goqu.I("create_at").Desc()).ToSQL()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	fmt.Println(strSql)

	list, err := DoQuery(u.Ctx, strSql, defaultDB, func() entity.Entity {
		return &user.User{}
	})
	if err != nil {
		return nil, nil
	}

	var userList []user.User
	for _, v := range list {
		if u2, ok := v.(*user.User); ok {
			userList = append(userList, *u2)
		}
	}

	return userList, nil
}

func (u *UserRepository) FindOne(id string) (*user.User, error) {
	userId, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	usr := &user.User{UserID: userId}
	if err := entity.Load(u.Ctx, usr, defaultDB); err != nil {
		return nil, errors.New(err.Error())
	}

	usr.SetPersistence()
	return usr, nil
}

func (u *UserRepository) FindByName(name string) (*user.User, error) {
	usr := &user.User{UserName: name}
	strSql, _, err := goqu.From(usr.TableName()).Where(goqu.Ex{"user_name": name}).ToSQL()

	if err != nil {
		return nil, errors.New(err.Error())
	}

	if err := DoQueryOne(u.Ctx, strSql, usr, defaultDB); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New(err.Error())
	}

	usr.SetPersistence()
	return usr, nil
}

func (u *UserRepository) Create(user *user.User) error {
	if _, err := entity.Insert(u.Ctx, user, defaultDB); err != nil {
		return errors.New(err.Error())
	} else {
		return nil
	}
}

func (u *UserRepository) Update(user *user.User) error {
	if err := entity.Update(u.Ctx, user, defaultDB); err != nil {
		return errors.New(err.Error())
	} else {
		return nil
	}
}
