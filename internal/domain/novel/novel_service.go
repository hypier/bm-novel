package novel

import (
	"bm-novel/internal/domain/novel/draft"
	"bm-novel/internal/http/web"
	"context"
	"io"
	"regexp"
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
	return s.Repo.Create(ctx, novel)
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

	dbNovel.ResponsibleEditorID = editorID
	return s.Repo.Update(ctx, dbNovel)

}

func findChapterLine(data []byte) (begin, end int) {
	p1 := `([零一二两三四五六七八九十百千万亿壹贰叁肆伍陆柒捌玖拾佰仟1-9]+)([集章回话节 、])([\w\W].*)`
	pos := regexp.MustCompile(p1).FindIndex(data)
	if len(pos) == 0 {
		return 0, 0
	}

	return pos[0], pos[1]
}

// UploadDraft 上传原文
func (s Service) UploadDraft(ctx context.Context, novelID uuid.UUID, file io.Reader) error {
	logrus.Debug("小说解析开始")
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Second)
	defer cancel()

	counter, err := s.Counter.FindOne(ctx, novelID)

	if err != nil {
		return err
	}

	d := draft.Draft{}
	d.Parser(counter, file)

	return nil
}
