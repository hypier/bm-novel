package novel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Chapter 章节
type Chapter struct {
	// 章节ID
	ChapterID uuid.UUID `json:"chapter_id" db:"chapter_id"`
	// 章节标题
	ChapterTitle string `json:"chapter_title" db:"chapter_title"`
	// 章节序号
	ChapterNo int `json:"chapter_no" db:"chapter_no"`
	// 标识 1 正确章 2 重复章 3 缺失章 4 错序章
	Flag int `json:"flag" db:"flag"`
	// 外包ID
	OutSourceID uuid.UUID `json:"out_source_id" db:"out_source_id"`
	// 字数
	WordsCount int `json:"words_count" db:"words_count"`
	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id"`

	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
}
