package base

import "time"

// Entity 基础
type Entity struct {
	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
	IsDelete bool      `json:"is_delete" db:"is_delete"`
}
