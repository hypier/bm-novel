package user

import (
	"bm-novel/internal/infrastructure/security"
	"github.com/pkg/errors"
)

// 用户基本信息
type User struct {
	// 用户id
	UserId string `json:"userId"`
	// 用户名
	UserName string `json:"userName"`
	// 用户状态
	UserLock bool `json:"stats"`
	// 密码
	Password string `json:"password"`
	// 角色代码
	RoleCode string `json:"roleCode"`
	// 姓名
	TrueName string `json:"trueName"`

	// 是否持久化，内部参数
	isPersistence bool
	// 首次密码
	FirstPassword bool

	repo IUserRepository
}

func (u *User) Get(userId string) (*User, error) {
	return u.repo.FindOne(userId)
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
