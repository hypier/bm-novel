package user

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/postgres"
	"context"

	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/joyparty/entity"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

// Repository 用户持久化
type Repository struct {
	db *sqlx.DB
}

// New 创建持久化对象
func New() *Repository {
	return &Repository{db: postgres.DefaultDB}
}

// FindList 查询用户列表
func (u *Repository) FindList(ctx context.Context, roleCode []string, realName string, pageIndex int, pageSize int) (user.Users, error) {

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

	logrus.WithFields(logrus.Fields{
		"SQL": strSQL,
	}).Debug("user findList")

	users := &user.Users{}
	err = u.db.SelectContext(ctx, users, strSQL, params...)

	return *users, err
}

// FindOne 根据ID查询
func (u *Repository) FindOne(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	usr := &user.User{UserID: userID}
	if err := entity.Load(ctx, usr, u.db); err != nil {
		return nil, errors.New(err.Error())
	}

	return usr, nil
}

// FindByName 根据用户名查询用户
func (u *Repository) FindByName(ctx context.Context, name string) (*user.User, error) {
	usr := user.User{}
	strSQL, params, err := goqu.From(usr.TableName()).Where(goqu.Ex{"user_name": name}).ToSQL()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	users := &user.Users{}
	err = u.db.SelectContext(ctx, users, strSQL, params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	for _, v := range *users {
		return v, nil
	}

	return nil, err
}

// Create 创建
func (u *Repository) Create(ctx context.Context, user *user.User) error {
	if _, err := entity.Insert(ctx, user, u.db); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// Update 更新
func (u *Repository) Update(ctx context.Context, user *user.User) error {
	if err := entity.Update(ctx, user, u.db); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
