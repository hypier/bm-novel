package chapter

import (
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/postgres"
	"context"
	"fmt"

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

func (r Repository) BatchCreate(ctx context.Context, cs *chapter.Chapters) error {
	c := &chapter.Chapter{}

	insert, err := entity.PrepareInsert(ctx, c, r.db)
	if err != nil {
		return err
	}

	for _, p2 := range *cs {
		_, err := insert.ExecContext(ctx, p2)
		if err != nil {
			fmt.Println("Exec error:", err)
			panic(err)
		}
	}

	return nil
}

func (r Repository) Update(ctx context.Context, chapter *chapter.Chapter) error {
	panic("implement me")
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
