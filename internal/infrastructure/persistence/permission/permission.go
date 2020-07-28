package permission

import (
	"bm-novel/internal/domain/permission"
	"bm-novel/internal/infrastructure/postgres"
	"context"
	"fmt"

	"github.com/joyparty/entity"

	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
)

// Repository 权限持久化
type Repository struct {
	Ctx context.Context
}

// FindAll 获取所有权限点
func (p *Repository) FindAll() (*permission.Permissions, error) {
	per := &permission.Permission{}
	strSQL, params, err := goqu.From(per.TableName()).ToSQL()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	fmt.Println(strSQL)

	permissions := &permission.Permissions{}
	err = postgres.DefaultDB.SelectContext(p.Ctx, permissions, strSQL, params...)

	return permissions, err
}

// Create 创建权限点
func (p *Repository) Create(permission *permission.Permission) error {
	if _, err := entity.Insert(p.Ctx, permission, postgres.DefaultDB); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

// BatchCreate 批量创建
func (p *Repository) BatchCreate(permissions *permission.Permissions) error {
	for _, v := range *permissions {
		_ = p.Create(v)
	}

	return nil
}
