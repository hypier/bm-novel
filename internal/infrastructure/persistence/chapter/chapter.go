package chapter

import (
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/postgres"
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/joyparty/entity"
	"github.com/sirupsen/logrus"
)

// Repository 权限持久化
type Repository struct {
	db *sqlx.DB
}

// New 创建持久化对象
func New() *Repository {
	return &Repository{db: postgres.DefaultDB}
}

// Create 创建
func (r Repository) Create(ctx context.Context, chapter *chapter.Chapter) error {
	if _, err := entity.Insert(ctx, chapter, r.db); err != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"chapter": &chapter,
		}, err, "Chapter Create exce SQL")
	}

	return nil
}
