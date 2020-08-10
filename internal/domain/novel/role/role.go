package role

import (
	"context"
	"time"

	"github.com/joyparty/entity"

	uuid "github.com/satori/go.uuid"
)

// Role 角色
type Role struct {
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

// TableName 表
func (r *Role) TableName() string {
	return "novel_role"
}

// OnEntityEvent 事件
func (r *Role) OnEntityEvent(ctx context.Context, ev entity.Event) error {
	switch ev {
	case entity.EventBeforeInsert:
		r.CreateAt = time.Now()
		r.UpdateAt = time.Now()
	case entity.EventBeforeUpdate:
		r.UpdateAt = time.Now()
	}

	return nil
}
