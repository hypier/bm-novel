package novel

import (
	"bm-novel/internal/domain/novel/chapter"
	nc "bm-novel/internal/domain/novel/counter"
	"bm-novel/internal/domain/novel/paragraph"
	"bm-novel/internal/domain/novel/role"
	"context"
	"io"

	uuid "github.com/satori/go.uuid"
)

// INovelService 小说服务接口
type INovelService interface {

	// 创建小说
	Create(ctx context.Context, novel *Novel) error
	// 删除小说
	Delete(ctx context.Context, novelID uuid.UUID) error
	// 指派责编
	AssignResponsibleEditor(ctx context.Context, novelID uuid.UUID, editorID uuid.UUID) error
	// 上传源文
	UploadDraft(ctx context.Context, novelID uuid.UUID, file io.Reader) error
}

// INovelRepository 小说仓库
type INovelRepository interface {
	FindList(ctx context.Context, novelName string, pageIndex int, pageSize int) (Novels, error)
	FindOne(ctx context.Context, novelID uuid.UUID) (*Novel, error)
	FindByTitle(ctx context.Context, title string) (*Novel, error)
	Create(ctx context.Context, novel *Novel) error
	Update(ctx context.Context, novel *Novel) error
}

// INovelCounterRepository 计数器
type INovelCounterRepository interface {
	FindOne(ctx context.Context, counterID uuid.UUID) (*nc.NovelCounter, error)
	FindByNovelID(ctx context.Context, novelID uuid.UUID) (*nc.NovelCounter, error)
	Create(ctx context.Context, counter *nc.NovelCounter) error
	Update(ctx context.Context, counter *nc.NovelCounter) error
}

// IChapterRepository 章节仓库
type IChapterRepository interface {
	Create(ctx context.Context, chapter *chapter.Chapter) error
	BatchCreate(ctx context.Context, chapters *chapter.Chapters) error
	Update(ctx context.Context, chapter *chapter.Chapter) error
}

// IParagraphRepository 段落仓库
type IParagraphRepository interface {
	Create(ctx context.Context, paragraph *paragraph.Paragraph) error
	BatchCreate(ctx context.Context, paragraphs *paragraph.Paragraphs) error
	Update(ctx context.Context, paragraph *paragraph.Paragraph) error
}

// IChapterService 章节服务
type IChapterService interface {
	// 保存章节
	Save(chapter *[]chapter.Chapter)
	// 编辑原文章节
	EditSource(source string) error
	// 删除章节
	Delete() error
	// 指派普通外包编辑
	AssignOrdinaryEditor(userID int) error
	// 段落解析
	ParseParagraph() (*[]paragraph.Paragraph, error)
}

// IParagraphService 段落
type IParagraphService interface {
	// 编辑段落
	Edit(source string) error
	// 删除段落
	delete() error
	// 指派角色
	AssignRole(roleCode string) error
	// 合并角色
	MergeRole(sourceRoleCode, targetRoleCode string) error
}

// IRoleService 角色服务
type IRoleService interface {
	// 创建角色
	Create(role *role.Role) error
	// 编辑角色
	Edit(role *role.Role) error
	// 删除角色
	Delete(roleCode string) error
}

// IEpisodeService 集数接口
type IEpisodeService interface {
	// 定集
	Create(paragraphCode *[]string)
	// 提审
	Apply()
	// 审核
	Audit()
	// 定稿
	Finalization()
}
