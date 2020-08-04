package novel

// INovelService 小说服务接口
type INovelService interface {

	// 创建小说
	Create(novel *Novel) error
	// 删除小说
	Delete() error
	// 指派责编
	AssignResponsibleEditor(userID string) error
	// 设置章节解析格式
	SetChapterFormat(format *Setting) error
	// 上传源文
	UploadSource(source string) error
	// 解析章节
	ParseChapters(format *Setting) (*[]Chapter, error)
}

// IChapterService 章节服务
type IChapterService interface {
	// 保存章节
	Save(chapter *[]Chapter)
	// 编辑原文章节
	EditSource(source string) error
	// 删除章节
	Delete() error
	// 指派普通外包编辑
	AssignOrdinaryEditor(userID int) error
	// 段落解析
	ParseParagraph() (*[]Paragraph, error)
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
	Create(role *NovelRole) error
	// 编辑角色
	Edit(role *NovelRole) error
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
