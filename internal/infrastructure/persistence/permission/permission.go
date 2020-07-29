package permission

import (
	"bm-novel/internal/domain/permission"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/postgres"
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/joyparty/entity"

	"github.com/doug-martin/goqu/v9"
)

// Repository 权限持久化
type Repository struct {
	db *sqlx.DB
}

// New 创建持久化对象
func New() *Repository {
	return &Repository{db: postgres.DefaultDB}
}

// FindAll 获取所有权限点
func (p *Repository) FindAll(ctx context.Context) (permission.Permissions, error) {
	per := &permission.Permission{}
	strSQL, params, err := goqu.From(per.TableName()).ToSQL()
	if err != nil {
		return nil, web.WriteErrLog(err, "Permission FindList goqu SQL")
	}

	permissions := &permission.Permissions{}
	err = p.db.SelectContext(ctx, permissions, strSQL, params...)
	if err != nil {
		err = web.WriteErrLog(err, "Permission FindList exce SQL")
	}
	return *permissions, err
}

// Create 创建权限点
func (p *Repository) Create(ctx context.Context, permission *permission.Permission) error {
	if _, err := entity.Insert(ctx, permission, p.db); err != nil {
		return web.WriteErrLog(err, "Permission insert exce SQL")
	}

	return nil
}

// BatchCreate 批量创建
func (p *Repository) BatchCreate(ctx context.Context, permissions *permission.Permissions) error {
	for _, v := range *permissions {
		_ = p.Create(ctx, v)
	}

	return nil
}
