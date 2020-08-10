package paragraph

import (
	"bm-novel/internal/domain/novel/base"
	"context"
	"time"

	"github.com/joyparty/entity"

	uuid "github.com/satori/go.uuid"
)

// Paragraphs 段落集
type Paragraphs []*Paragraph

// Paragraph 段落
type Paragraph struct {
	// 段落ID
	ParagraphID uuid.UUID `json:"paragraph_id" db:"paragraph_id,primaryKey"`
	// 章节索引
	ChapterIndex int `json:"chapter_index" db:"chapter_index"`
	// 集索引
	EpisodeIndex int `json:"episode_index" db:"episode_index"`
	// 内容
	Content string `json:"content" db:"content"`
	// 字数
	WordsCount int `json:"words_count" db:"words_count"`

	// 章节ID
	ChapterID uuid.NullUUID `json:"chapter_id" db:"chapter_id"`
	// 集数ID
	EpisodeID uuid.NullUUID `json:"episode_id" db:"episode_id"`
	// 角色ID
	NovelRoleID uuid.NullUUID `json:"novel_role_id" db:"role_id"`

	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id"`

	base.Entity
}

// TableName 表
func (p *Paragraph) TableName() string {
	return "paragraph"
}

// OnEntityEvent 事件
func (p *Paragraph) OnEntityEvent(ctx context.Context, ev entity.Event) error {
	switch ev {
	case entity.EventBeforeInsert:
		p.CreateAt = time.Now()
		p.UpdateAt = time.Now()
	case entity.EventBeforeUpdate:
		p.UpdateAt = time.Now()
	}

	return nil
}
