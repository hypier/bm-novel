package novel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Episode 集数
type Episode struct {
	// 集数ID
	EpisodeID uuid.UUID `json:"episode_id" db:"episode_id"`
	// 集数标题
	EpisodeTitle string `json:"episode_title" db:"episode_title"`
	// 集数序号
	EpisodeNo string `json:"episode_no" db:"episode_no"`
	// 状态  1 未审核 2审核中 4已审核 8已定稿
	Status int `json:"status" db:"status"`

	// 字数
	WordsCount int `json:"words_count" db:"words_count"`
	// 小说ID
	NovelID uuid.UUID `json:"novel_id" db:"novel_id"`

	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
}
