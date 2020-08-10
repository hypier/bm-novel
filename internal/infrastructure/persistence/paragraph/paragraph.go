package paragraph

import (
	"bm-novel/internal/domain/novel/paragraph"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/postgres"
	"bm-novel/internal/infrastructure/redis"
	"context"
	"encoding/json"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/doug-martin/goqu/v9"

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

// Update 更新
func (r Repository) Update(ctx context.Context, paragraph *paragraph.Paragraph) error {
	panic("implement me")
}

// Create 创建
func (r Repository) Create(ctx context.Context, paragraph *paragraph.Paragraph) error {
	if _, err := entity.Insert(ctx, paragraph, r.db); err != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"paragraph": &paragraph,
		}, err, "Paragraph Create exce SQL")
	}

	return nil
}

// BatchCreate 批量创建
func (r Repository) BatchCreate(ctx context.Context, ps *paragraph.Paragraphs) error {
	return r.saveDB(ctx, ps)
}

func (r Repository) saveRedis(ctx context.Context, ps *paragraph.Paragraphs) error {

	var values []interface{}
	var novelID uuid.UUID
	for i, p := range *ps {
		if i == 0 {
			novelID = p.NovelID
		}

		marshal, err := json.Marshal(p)
		if err != nil {
			logrus.Warnf("putCache marshal marshal err, %s", err)
			continue
		}

		field := p.Index
		values = append(values, field)
		values = append(values, marshal)
	}

	key := fmt.Sprintf("novel:paragraph:%s", novelID.String())
	redis.GetChcher().HMPut(key, time.Hour, values)
	return nil
}

func (r Repository) saveDB(ctx context.Context, ps *paragraph.Paragraphs) error {
	p := &paragraph.Paragraph{}

	insert, err := entity.PrepareInsert(ctx, p, r.db)
	if err != nil {
		return err
	}

	for _, p2 := range *ps {
		_, err := insert.ExecContext(ctx, p2)
		if err != nil {
			fmt.Println("Exec error:", err)
			panic(err)
		}
	}
	return nil
}

// FindAll 获取所有权限点
func (r *Repository) FindAll(ctx context.Context) (paragraph.Paragraphs, error) {

	par := &paragraph.Paragraph{}
	strSQL, params, err := goqu.From(par.TableName()).ToSQL()
	if err != nil {
		return nil, web.WriteErrLog(err, "Paragraph FindList goqu SQL")
	}

	pars := &paragraph.Paragraphs{}
	err = r.db.SelectContext(ctx, pars, strSQL, params...)
	if err != nil {
		err = web.WriteErrLog(err, "Paragraph FindList exce SQL")
	}
	return *pars, err
}

// FindList 查询用户列表
func (r *Repository) FindList(ctx context.Context, prevIndex int, pageSize int) (*paragraph.Paragraphs, error) {

	par := &paragraph.Paragraph{}

	strSQL, params, err := goqu.From(par.TableName()).
		Where(goqu.C("index").Gte(prevIndex)).
		Limit(uint(pageSize)).
		Order(goqu.I("index").Asc(), goqu.I("sub_index").Desc()).ToSQL()
	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"SQL": strSQL,
	}).Debug("user findList")

	pars := &paragraph.Paragraphs{}
	err = r.db.SelectContext(ctx, pars, strSQL, params...)

	if err != nil {
		return nil, err
	}

	return pars, err
}
