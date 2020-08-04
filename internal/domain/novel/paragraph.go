package novel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Paragraph 段落
type Paragraph struct {
	// 段落ID
	ParagraphID uuid.UUID `json:"paragraph_id" db:"paragraph_id"`
	// 内容
	Content string `json:"content" db:"content"`
	// 下一个段落
	Next uuid.UUID `json:"next" db:"next"`
	// 上一个段落
	Prev uuid.UUID `json:"prev" db:"prev"`
	// 字数
	WordsCount int `json:"words_count" db:"words_count"`

	// 章节ID
	ChapterID uuid.UUID `json:"chapter_id" db:"chapter_id"`
	// 集数ID
	EpisodeID uuid.UUID `json:"episode_id" db:"episode_id"`
	// 角色ID
	NovelRoleID uuid.UUID `json:"novel_role_id" db:"novel_role_id"`

	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id"`

	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
}
