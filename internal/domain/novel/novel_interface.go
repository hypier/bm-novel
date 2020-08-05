package novel

import (
	"bm-novel/internal/domain/novel/chapter"
	"bm-novel/internal/domain/novel/paragraph"
	"bm-novel/internal/domain/novel/role"
	"context"

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
	UploadDraft(ctx context.Context, novelID uuid.UUID, draft string) error
}

type INovelRepository interface {
	FindList(ctx context.Context, novelName string, pageIndex int, pageSize int) (Novels, error)
	FindOne(ctx context.Context, novelID uuid.UUID) (*Novel, error)
	FindByTitle(ctx context.Context, title string) (*Novel, error)
	Create(ctx context.Context, novel *Novel) error
	Update(ctx context.Context, novel *Novel) error
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
