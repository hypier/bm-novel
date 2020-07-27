package permission

import (
	"context"

	"github.com/joyparty/entity"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Permissions []*Permission

type Permission struct {
	PID    uuid.UUID      `json:"pid" db:"pid,primaryKey"`
	Name   string         `json:"name" db:"name"`
	URI    string         `json:"uri" db:"uri"`
	Method string         `json:"method" db:"method"`
	Roles  pq.StringArray `json:"roles" db:"roles"`
}

// TableName 表名
func (p Permission) TableName() string {
	return "permission"
}

// OnEntityEvent 持久化事件
func (p Permission) OnEntityEvent(ctx context.Context, ev entity.Event) error {
	return nil
}
