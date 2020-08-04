package novel

import (
	"bm-novel/internal/http/web"
	"context"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

// Service 小说服务
type Service struct {
	Repo INovelRepository
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

	return s.Repo.Create(ctx, novel)
}

// Delete 删除小说
func (s Service) Delete(ctx context.Context, novelID uuid.UUID) error {
	panic("implement me")
}

// AssignResponsibleEditor 指派责编
func (s Service) AssignResponsibleEditor(ctx context.Context, novelID uuid.UUID, editorID uuid.UUID) error {
	panic("implement me")
}

// SetFormat 设置格式
func (s Service) SetFormat(ctx context.Context, novelID uuid.UUID, format Settings) error {
	panic("implement me")
}

// UploadDraft 上传原文
func (s Service) UploadDraft(ctx context.Context, novelID uuid.UUID, draft string) error {
	panic("implement me")
}
