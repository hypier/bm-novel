package counter

import (
	"bm-novel/internal/domain/novel/base"
	"context"

	"github.com/joyparty/entity"
	uuid "github.com/satori/go.uuid"
)

// NovelCounters 计数器集合
type NovelCounters []*NovelCounter

// NovelCounter 小说计数器
type NovelCounter struct {
	CountID uuid.UUID `json:"count_id" db:"count_id,primaryKey"`
	// 已指派的章节数
	AssignedChaptersCount int `json:"assigned_chapters_count" db:"assigned_chapters_count"`
	// 总章节数
	ChaptersCount int `json:"chapters_count" db:"chapters_count"`
	// 总字数
	WordsCount int `json:"words_count" db:"words_count"`
	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id"`

	base.Entity
}

// TableName 表名
func (c *NovelCounter) TableName() string {

	return "novel_counter"
}

// OnEntityEvent 事件
func (c *NovelCounter) OnEntityEvent(ctx context.Context, ev entity.Event) error {

	return nil
}
