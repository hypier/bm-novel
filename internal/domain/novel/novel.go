package novel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Novel 小说
type Novel struct {
	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id,primaryKey"`
	// 已指派的章节数
	AssignedChaptersCount int `json:"assigned_chapters_count"`
	// 总章节数
	ChaptersCount int `json:"chapters_count"`
	// 主编ID
	ChiefEditorID uuid.UUID `json:"chief_editor_id" db:"chief_editor_id"`
	// 小说标题
	NovelTitle string `json:"novel_title" db:"novel_title"`
	// 责编ID
	ResponsibleEditorID uuid.UUID `json:"responsible_editor_id" db:"responsible_editor_id"`
	// 总字数
	WordsCount int `json:"words_count" db:"words_count"`

	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`

	Settings string
}

// Settings 章节格式
type Settings struct {
	// 章节是否有前缀
	HasPrefix bool `json:"has_prefix"`
	// 章节前缀
	Prefix string `json:"prefix"`
	// 章节是否有后缀
	HasSuffix bool `json:"has_suffix"`
	// 章节后缀
	Suffix string `json:"suffix"`
	// 章节数格式 1 阿拉伯数字 2 中文 小写 3 中文大写
	Format int `json:"format"`
	// 单集最大字数
	WordsMax int `json:"words_max"`
	// 单集最小字数
	WordsMin int `json:"words_min"`
	// 章节分隔符 1 换行符 2 空格
	Separator int `json:"separator"`
}
