package novel

import (
	"bm-novel/internal/domain/novel"
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

// FindList 查询
func (r Repository) FindList(ctx context.Context, novelName string, pageIndex int, pageSize int) (novel.Novels, error) {

	panic("implement me")
}

// FindOne 单个查询
func (r Repository) FindOne(ctx context.Context, novelID uuid.UUID) (*novel.Novel, error) {
	n := &novel.Novel{NovelID: novelID}
	if err := entity.Load(ctx, n, r.db); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 数据为空
			return nil, nil
		}

		return nil, web.WriteErrLogWithField(logrus.Fields{
			"novelID": novelID,
		}, err, "Novel FindOne exce SQL")
	}

	return n, nil
}

// FindByTitle 当标题查询
func (r Repository) FindByTitle(ctx context.Context, title string) (*novel.Novel, error) {
	n := &novel.Novel{}

	strSQL, params, err := goqu.From(n.TableName()).Where(goqu.Ex{"novel_title": title}).ToSQL()
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"sql": strSQL,
	}).Debug("Repository FindByTitle")

	novels := &novel.Novels{}
	err = r.db.SelectContext(ctx, novels, strSQL, params...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 数据为空
			return nil, nil
		}

		return nil, web.WriteErrLogWithField(logrus.Fields{
			"strSQL": strSQL,
			"title":  title,
		}, err, "Novel FindTitle exce SQL")
	}

	for _, v := range *novels {
		return v, nil
	}

	return nil, err
}

// Create 创建
func (r Repository) Create(ctx context.Context, novel *novel.Novel) error {
	if _, err := entity.Insert(ctx, novel, r.db); err != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"novel": &novel,
		}, err, "Novel Create exce SQL")
	}

	return nil
}

// Update 更新
func (r Repository) Update(ctx context.Context, novel *novel.Novel) error {
	if err := entity.Update(ctx, novel, r.db); err != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"novel": &novel,
		}, err, "Novel Update exce SQL")
	}

	return nil
}
