package domain

import "bm-story/internal/domain/story"

type StoryService interface {

	// 创建小说
	Create(story *story.Story) error
	// 删除小说
	Delete() error
	// 指派责编
	AssignResponsibleEditor(userId int) error
	// 设置章节解析格式
	SetChapterFormat(format *story.ChapterFormat) error
	// 上传源文
	UploadSource(source string) error
	// 解析章节
	ParseChapters(format *story.ChapterFormat) (*[]story.Chapter, error)
}

type ChapterService interface {
	// 保存章节
	Save(chapter *[]story.Chapter)
	// 编辑原文章节
	EditSource(source string) error
	// 删除章节
	Delete() error
	// 指派普通外包编辑
	AssignOrdinaryEditor(userId int) error
	// 段落解析
	ParseParagraph() (*[]story.Paragraph, error)
}

type ParagraphService interface {
	// 编辑段落
	Edit(source string) error
	// 删除段落
	delete() error
	// 指派角色
	AssignRole(roleCode string) error
	// 合并角色
	MergeRole(sourceRoleCode, targetRoleCode string) error
}

type RoleService interface {
	// 创建角色
	Create(role *story.Role) error
	// 编辑角色
	Edit(role *story.Role) error
	// 删除角色
	Delete(roleCode string) error
}

type EpisodeService interface {
	// 定集
	Create(paragraphCode *[]string)
	// 提审
	Apply()
	// 审核
	Audit()
	// 定稿
	Finalization()
}
