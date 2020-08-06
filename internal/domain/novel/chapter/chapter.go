package chapter

import (
	"context"
	"time"

	"github.com/joyparty/entity"

	uuid "github.com/satori/go.uuid"
)

type Chapters []*Chapter

// Chapter 章节
type Chapter struct {
	// 章节ID
	ChapterID uuid.UUID `json:"chapter_id" db:"chapter_id,primaryKey"`
	// 章节标题
	ChapterTitle string `json:"chapter_title" db:"chapter_title"`
	// 章节卷号
	Volume int `json:"volume" db:"volume"`
	// 章节序号
	ChapterNo int `json:"chapter_no" db:"chapter_no"`
	// 标识 1 正确章 2 重复章 3 缺失章 4 错序章
	FeatureCode int `json:"feature_code" db:"feature_code"`
	// 外包ID
	OutSourceID uuid.UUID `json:"out_source_id" db:"out_source_id"`
	// 字数
	WordsCount int `json:"words_count" db:"words_count"`
	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id"`

	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
}

func (c Chapter) TableName() string {

	return "chapter"
}

func (c Chapter) OnEntityEvent(ctx context.Context, ev entity.Event) error {
	switch ev {
	case entity.EventBeforeInsert:
		c.CreateAt = time.Now()
		c.UpdateAt = time.Now()
	case entity.EventBeforeUpdate:
		c.UpdateAt = time.Now()
	}

	return nil
}
