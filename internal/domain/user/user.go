package user

import (
	"context"
	"time"

	"github.com/joyparty/entity"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// Users 会员数组
type Users []*User

// User 用户基本信息.
type User struct {
	// 用户id
	UserID uuid.UUID `json:"user_id" db:"user_id,primaryKey"`
	// 用户名
	UserName string `json:"user_name" db:"user_name"`
	// 用户状态
	IsLock bool `json:"is_lock" db:"is_lock"`
	// 密码
	Password string `json:"password" db:"password"`
	// 角色代码
	Roles pq.StringArray `json:"roles" db:"roles"`
	// 姓名
	RealName string `json:"real_name" db:"real_name"`
	// 是否需要修改密码
	NeedChangePassword bool `json:"need_change_password" db:"need_change_password"`

	CreateAt time.Time `db:"create_at,refuseUpdate"`
	UpdateAt time.Time `db:"update_at"`
}

// TableName is db table
func (u *User) TableName() string {
	return "user"
}

// OnEntityEvent 存储事件回调方法，entity.Entity接口方法
func (u *User) OnEntityEvent(ctx context.Context, ev entity.Event) error {
	switch ev {
	case entity.EventBeforeInsert:
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
	case entity.EventBeforeUpdate:
		u.UpdateAt = time.Now()
	}

	return nil
}
