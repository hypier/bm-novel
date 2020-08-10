package novel

import (
	nc "bm-novel/internal/domain/novel/counter"
	"bm-novel/internal/domain/novel/draft"
	"bm-novel/internal/http/web"
	"context"
	"io"
	"time"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

// Service 小说服务
type Service struct {
	Repo          INovelRepository
	Counter       INovelCounterRepository
	ChapterRepo   IChapterRepository
	ParagraphRepo IParagraphRepository
}

// Create 创建小说
func (s Service) Create(ctx context.Context, novel *Novel) error {
	title := novel.NovelTitle

	dbNovel, err := s.Repo.FindByTitle(ctx, title)
	if err != nil {
		return err
	}

	if dbNovel != nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"title":     title,
			"dbNovelID": dbNovel.NovelID,
		}, web.ErrConflict, "Create Novel, Duplicate novelTitle")
	}

	novel.NovelID = uuid.NewV4()
	err = s.Repo.Create(ctx, novel)
	if err != nil {
		return err
	}

	counter := &nc.NovelCounter{CountID: uuid.NewV4(), NovelID: novel.NovelID}
	return s.Counter.Create(ctx, counter)
}

// Delete 删除小说
func (s Service) Delete(ctx context.Context, novelID uuid.UUID) error {
	panic("implement me")
}

// AssignResponsibleEditor 指派责编
func (s Service) AssignResponsibleEditor(ctx context.Context, novelID uuid.UUID, editorID uuid.UUID) error {
	dbNovel, err := s.Repo.FindOne(ctx, novelID)

	if err != nil {
		return err
	}

	if dbNovel == nil {
		return web.WriteErrLogWithField(logrus.Fields{
			"novelID": novelID,
		}, web.ErrNotFound, "AssignResponsibleEditor, Novel Not Found")
	}

	dbNovel.ResponsibleEditorID = uuid.NullUUID{UUID: editorID}
	return s.Repo.Update(ctx, dbNovel)

}

// UploadDraft 上传原文
func (s Service) UploadDraft(ctx context.Context, novelID uuid.UUID, file io.Reader) error {
	logrus.Debug("小说解析开始")
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	counter, err := s.Counter.FindByNovelID(ctx, novelID)

	if err != nil {
		return err
	}

	if counter == nil {
		return web.ErrNotFound
	}

	d := draft.Draft{}
	d.Parser(counter, file)
	logrus.Debug("小说解析结束")
	if err = s.ChapterRepo.BatchCreate(ctx, &d.Chapters); err != nil {
		return err
	}

	if err = s.ParagraphRepo.BatchCreate(ctx, d.Paragraphs); err != nil {
		return err
	}

	if err = s.Counter.Update(ctx, d.Counter); err != nil {
		return err
	}

	logrus.Debug("小说保存结束")
	return nil
}
