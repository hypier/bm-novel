package user

import (
	"bm-novel/internal/infrastructure/security"
	"context"
	"github.com/joyparty/entity"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

// 用户基本信息
type User struct {
	// 用户id
	UserId string `json:"userId" db:"user_id,primaryKey"`
	// 用户名
	UserName string `json:"userName" db:"user_name"`
	// 用户状态
	IsLock bool `json:"isLock" db:"is_lock"`
	// 密码
	Password string `json:"password" db:"password"`
	// 角色代码
	RoleCode pq.StringArray `json:"roleCode" db:"role_code"`
	// 姓名
	RealName string `json:"realName" db:"real_name"`
	// 是否需要重设密码
	NeedResetPassword bool `json:"needResetPassword" db:"need_reset_password"`

	// 是否持久化，内部参数
	isPersistence bool      `db:"-"`
	CreateAt      time.Time `db:"create_at,refuseUpdate"`
	UpdateAt      time.Time `db:"update_at"`

	Repo IUserRepository `db:"-"`
}

func (u *User) Create(user *User) error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	u.UserId = uuid.NewV4().String()

	u.NeedResetPassword = true
	return u.Repo.Create(user)
}

func (u *User) SetRole(roleCode []string) error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	u.RoleCode = roleCode
	return u.Repo.Update(u)
}

func (u *User) InitPassword(password string) error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	if !u.NeedResetPassword {
		return errors.New("不能重复初始化密码")
	}

	hashPassword, err := security.Hash(password)
	if err != nil {
		return errors.New(err.Error())
	}
	u.Password = string(hashPassword)
	u.NeedResetPassword = false
	return u.Repo.Update(u)
}

func (u *User) ResetPassword() error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	hashPassword, err := security.Hash("123456")
	if err != nil {
		return errors.New(err.Error())
	}
	u.Password = string(hashPassword)
	u.NeedResetPassword = true

	return u.Repo.Update(u)
}

func (u *User) Lock() error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	if u.IsLock {
		return errors.New("此用户已锁定")
	}

	u.IsLock = true
	return u.Repo.Update(u)
}

func (u *User) Unlock() error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	if !u.IsLock {
		return errors.New("此用户未锁定")
	}

	u.IsLock = false
	return u.Repo.Update(u)
}

func (u *User) CheckPassword(password string) (bool, error) {
	if !u.isPersistence {
		return false, errors.New("没有持久化对象")
	}

	err := security.VerifyPassword(u.Password, password)
	if err != nil {
		return false, errors.New(err.Error())
	}

	return true, nil
}

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
