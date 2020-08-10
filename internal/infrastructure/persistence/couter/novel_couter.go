package couter

import (
	nc "bm-novel/internal/domain/novel/counter"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/postgres"
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"

	"github.com/jmoiron/sqlx"
	"github.com/joyparty/entity"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
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

// FindByNovelID 通过novelID查询
func (r *Repository) FindByNovelID(ctx context.Context, novelID uuid.UUID) (*nc.NovelCounter, error) {
	c := nc.NovelCounter{}

	strSQL, params, err := goqu.From(c.TableName()).Where(goqu.Ex{"novel_id": novelID}).ToSQL()
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"sql": strSQL,
	}).Debug("Repository FindByNovelID")

	users := &nc.NovelCounters{}
	err = r.db.SelectContext(ctx, users, strSQL, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 数据为空
			return nil, nil
		}

		return nil, web.WriteErrLogWithField(logrus.Fields{
			"strSQL":  strSQL,
			"novelID": novelID,
		}, err, "Counter FindByNovelID exce SQL")
	}

	for _, v := range *users {
		return v, nil
	}

	return nil, err
}

// FindOne 查询
func (r *Repository) FindOne(ctx context.Context, novelID uuid.UUID) (*nc.NovelCounter, error) {
	c := &nc.NovelCounter{NovelID: novelID}
	if err := entity.Load(ctx, c, r.db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 数据为空
			return nil, nil
		}

		return nil, web.WriteErrLogWithField(logrus.Fields{
			"novelID": novelID,
		}, err, "NovelCounter FindOne exce SQL")
	}

	return c, nil
}

// Create 创建
func (r *Repository) Create(ctx context.Context, counter *nc.NovelCounter) error {
	if _, err := entity.Insert(ctx, counter, r.db); err != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"counter": &counter,
		}, err, "NovelCounter Create exce SQL")
	}

	return nil
}

// Update 更新
func (r *Repository) Update(ctx context.Context, counter *nc.NovelCounter) error {
	if err := entity.Update(ctx, counter, r.db); err != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"counter": &counter,
		}, err, "NovelCounter Update exce SQL")
	}

	return nil
}
