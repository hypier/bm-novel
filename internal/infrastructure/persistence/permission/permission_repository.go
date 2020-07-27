package permission

import (
	"bm-novel/internal/domain/permission"
	"bm-novel/internal/infrastructure/persistence"
	"context"
	"fmt"

	"github.com/joyparty/entity"

	"github.com/doug-martin/goqu/v9"
	"github.com/pkg/errors"
)

// PermissionRepository 权限持久化
type PermissionRepository struct {
	Ctx context.Context
}

// FindAll 获取所有权限点
func (p PermissionRepository) FindAll() (*permission.Permissions, error) {
	per := &permission.Permission{}
	strSQL, params, err := goqu.From(per.TableName()).ToSQL()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	fmt.Println(strSQL)

	pers := &permission.Permissions{}
	err = persistence.DefaultDB.SelectContext(p.Ctx, pers, strSQL, params...)

	return pers, nil
}

// Create 创建权限点
func (p PermissionRepository) Create(permission *permission.Permission) error {
	if _, err := entity.Insert(p.Ctx, permission, persistence.DefaultDB); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
