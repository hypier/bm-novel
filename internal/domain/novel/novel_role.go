package novel

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// NovelRole 角色
type NovelRole struct {
	// 角色ID
	RoleID uuid.UUID `json:"role_id" db:"role_id"`
	// 年纪
	Age string `json:"age" db:"age"`
	// 人设
	Characters string `json:"characters" db:"characters"`
	// 性别
	Gender string `json:"gender" db:"gender"`
	// 角色类别
	RoleClass string `json:"role_class" db:"role_class"`
	// 角色名
	RoleName string `json:"role_name" db:"role_name"`

	CreateAt time.Time `json:"create_at" db:"create_at"`
	UpdateAt time.Time `json:"update_at" db:"update_at"`
}
