package chapter

import (
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/postgres"
	"bm-novel/internal/infrastructure/redis"
	"context"
	"encoding/json"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

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
	return r.saveRedis(ctx, cs)
}

func (r Repository) saveRedis(ctx context.Context, cs *chapter.Chapters) error {

	var values []interface{}
	var novelID uuid.UUID
	for i, p := range *cs {
		if i == 0 {
			novelID = p.NovelID
		}

		marshal, err := json.Marshal(p)
		if err != nil {
			logrus.Warnf("putCache marshal marshal err, %s", err)
			continue
		}

		field := p.ChapterID.String()
		values = append(values, field)
		values = append(values, marshal)
	}

	key := fmt.Sprintf("novel:chapter:%s", novelID.String())
	redis.GetChcher().HMPut(key, time.Hour, values)
	return nil
}

func (r Repository) saveDB(ctx context.Context, cs *chapter.Chapters) error {
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
