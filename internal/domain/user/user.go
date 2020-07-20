package user

import (
	"bm-novel/internal/infrastructure/security"
	"context"
	"github.com/joyparty/entity"
	"github.com/pkg/errors"
	"time"
)

// 用户基本信息
type User struct {
	// 用户id
	UserId string `json:"userId" db:"user_id,primaryKey"`
	// 用户名
	UserName string `json:"userName" db:"user_name"`
	// 用户状态
	UserLock bool `json:"userLock" db:"user_lock"`
	// 密码
	Password string `json:"password" db:"password"`
	// 角色代码
	RoleCode string `json:"roleCode" db:"role_code"`
	// 姓名
	TrueName string `json:"trueName" db:"true_name"`
	// 首次密码
	FirstPassword bool `json:"firstPassword" db:"first_password"`

	// 是否持久化，内部参数
	isPersistence bool  `db:"-"`
	CreateAt      int64 `db:"create_at,refuseUpdate"`
	UpdateAt      int64 `db:"update_at"`

	repo IUserRepository
}

func (u *User) Create(user *User) error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)

	u.FirstPassword = true
	return u.repo.Create(user)
}

func (u *User) SetRole(roleCode string) error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	u.RoleCode = roleCode
	return u.repo.Update(u)
}

func (u *User) InitPassword(password string) error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	if !u.FirstPassword {
		return errors.New("不能重复初始化密码")
	}

	hashPassword, err := security.Hash(password)
	if err != nil {
		return errors.New(err.Error())
	}
	u.Password = string(hashPassword)
	u.FirstPassword = false
	return u.repo.Update(u)
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

	return u.repo.Update(u)
}

func (u *User) Lock() error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	if u.UserLock {
		return errors.New("此用户已锁定")
	}

	u.UserLock = true
	return u.repo.Update(u)
}

func (u *User) Unlock() error {
	if !u.isPersistence {
		return errors.New("没有持久化对象")
	}

	if !u.UserLock {
		return errors.New("此用户未锁定")
	}

	u.UserLock = false
	return u.repo.Update(u)
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
	return "db_user"
}

// OnEntityEvent 存储事件回调方法，entity.Entity接口方法
func (u *User) OnEntityEvent(ctx context.Context, ev entity.Event) error {
	switch ev {
	case entity.EventBeforeInsert:
		u.CreateAt = time.Now().Unix()
	case entity.EventBeforeUpdate:
		u.UpdateAt = time.Now().Unix()
	}

	return nil
}
