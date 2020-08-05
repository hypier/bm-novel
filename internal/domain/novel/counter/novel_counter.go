package counter

import (
	"context"

	"github.com/joyparty/entity"
	uuid "github.com/satori/go.uuid"
)

// Counter 小说计数器
type Counter struct {
	CountID uuid.UUID `json:"count_id" db:"count_id"`
	// 已指派的章节数
	AssignedChaptersCount int `json:"assigned_chapters_count" db:"assigned_chapters_count"`
	// 总章节数
	ChaptersCount int `json:"chapters_count" db:"chapters_count"`
	// 总字数
	WordsCount int `json:"words_count" db:"words_count"`
	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id"`
}

// TableName 表名
func (c Counter) TableName() string {

	return "novel_counter"
}

// OnEntityEvent 事件
func (c Counter) OnEntityEvent(ctx context.Context, ev entity.Event) error {
	panic("implement me")
}
