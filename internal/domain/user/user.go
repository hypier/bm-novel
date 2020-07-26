package user

import (
	"bm-novel/internal/infrastructure/security"
	"context"
	"time"

	"github.com/joyparty/entity"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var (
	// ErrUserConflict 用户名重复错误.
	ErrUserConflict = errors.New("User Conflict")
	// ErrUserNotFound 用户不存在.
	ErrUserNotFound = errors.New("User Not Found")
	// ErrUserLocked 用户被锁定.
	ErrUserLocked = errors.New("User Locked")
	// ErrNotAcceptable 不接受修改.
	ErrNotAcceptable = errors.New("Not Acceptable")
	// ErrPasswordIncorrect 用户名或密码错误.
	ErrPasswordIncorrect = errors.New("username or password is incorrect")
	// DefaultPassword 默认密码.
	DefaultPassword = "123456"
)

// Users 会员数组
type Users []*User

// User 用户基本信息.
type User struct {
	// 用户id
	UserID uuid.UUID `json:"userID" db:"user_id,primaryKey"`
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
	// 是否需要修改密码
	NeedChangePassword bool `json:"needChangePassword" db:"need_change_password"`

	CreateAt time.Time `db:"create_at,refuseUpdate"`
	UpdateAt time.Time `db:"update_at"`

	// 是否持久化，内部参数
	isPersistence bool
	// 持久化对象
	repo IUserRepository
}

// New 创建一个带持久化的对象
func New(repo IUserRepository) IUserServer {
	return &User{repo: repo}
}

// SetPersistence 是否被持久化到数据库
func (u *User) SetPersistence() {
	u.isPersistence = true
}

// SetRepo 设置持久化对象
func (u *User) SetRepo(repo IUserRepository) {
	u.repo = repo
}

// Load 载入对象
func (u *User) Load(userID string) (*User, error) {
	repo := u.repo

	u, err := u.repo.FindOne(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	u.repo = repo
	return u, nil
}

// Create 创建用户
func (u *User) Create(user User) (*User, error) {
	hashPassword, err := security.Hash(DefaultPassword)
	if err != nil {
		return nil, err
	}

	dbUser, err := u.repo.FindByName(user.UserName)

	if err != nil {
		return nil, err
	} else if dbUser != nil && dbUser.UserName == user.UserName {
		return nil, ErrUserConflict
	}

	u.Password = string(hashPassword)
	u.UserID = uuid.NewV4()
	u.NeedChangePassword = true
	u.UserName = user.UserName
	u.RoleCode = user.RoleCode
	u.RealName = user.RealName

	return u, u.repo.Create(u)
}

// Edit 编辑
func (u *User) Edit(user User) error {
	if !u.isPersistence {
		return ErrUserNotFound
	}

	u.RealName = user.RealName
	u.RoleCode = user.RoleCode
	u.UserName = user.UserName

	// 查询是否重复
	dbUser, err := u.repo.FindByName(user.UserName)
	if err != nil {
		return errors.New(err.Error())
	} else if dbUser != nil && dbUser.UserID != u.UserID {
		return ErrUserConflict
	}

	return u.repo.Update(u)
}

// ChangeInitPassword 修改初始密码
func (u *User) ChangeInitPassword(password string) error {
	if !u.isPersistence {
		return ErrUserNotFound
	}

	if !u.NeedChangePassword {
		return errors.New("不能重复初始化密码")
	}

	hashPassword, err := security.Hash(password)
	if err != nil {
		return errors.New(err.Error())
	}
	u.Password = string(hashPassword)
	u.NeedChangePassword = false
	return u.repo.Update(u)
}

// ResetPassword 重置密码
func (u *User) ResetPassword() error {
	if !u.isPersistence {
		return ErrUserNotFound
	}

	hashPassword, err := security.Hash(DefaultPassword)
	if err != nil {
		return errors.New(err.Error())
	}
	u.Password = string(hashPassword)
	u.NeedChangePassword = true

	return u.repo.Update(u)
}

// Lock 锁定
func (u *User) Lock() error {
	if !u.isPersistence {
		return ErrUserNotFound
	}

	if u.IsLock {
		return ErrNotAcceptable
	}

	u.IsLock = true
	return u.repo.Update(u)
}

// Unlock 解锁
func (u *User) Unlock() error {
	if !u.isPersistence {
		return ErrUserNotFound
	}

	if !u.IsLock {
		return ErrNotAcceptable
	}

	u.IsLock = false
	return u.repo.Update(u)
}

// CheckPassword ： 验证密码
func (u *User) CheckPassword(password string) error {
	if !u.isPersistence {
		return ErrUserNotFound
	}

	err := security.VerifyPassword(u.Password, password)
	if err != nil {
		return ErrPasswordIncorrect
	}

	return nil
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
