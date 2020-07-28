package user

import (
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/security"
	"context"

	uuid "github.com/satori/go.uuid"
)

var (
	// DefaultPassword 默认密码.
	DefaultPassword = "123456"
)

// Service user服务
type Service struct {
	Repo IUserRepository
}

// Create 创建用户
func (s Service) Create(ctx context.Context, user User) (*User, error) {
	hashPassword, err := security.Hash(DefaultPassword)
	if err != nil {
		return nil, err
	}

	dbUser, err := s.Repo.FindByName(ctx, user.UserName)

	if err != nil {
		return nil, err
	} else if dbUser != nil && dbUser.UserName == user.UserName {
		return nil, web.ErrUserConflict
	}

	u := &User{}
	u.Password = string(hashPassword)
	u.UserID = uuid.NewV4()
	u.NeedChangePassword = true
	u.UserName = user.UserName
	u.RoleCode = user.RoleCode
	u.RealName = user.RealName

	err = s.Repo.Update(ctx, u)

	return u, err
}

// Edit 编辑
func (s Service) Edit(ctx context.Context, userID uuid.UUID, user User) error {

	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	} else if dbUser != nil && dbUser.UserID != userID {
		return web.ErrUserConflict
	}

	dbUser.RealName = user.RealName
	dbUser.RoleCode = user.RoleCode
	dbUser.UserName = user.UserName

	return s.Repo.Update(ctx, dbUser)
}

// ChangeInitPassword 修改初始密码
func (s Service) ChangeInitPassword(ctx context.Context, userID uuid.UUID, password string) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	if !dbUser.NeedChangePassword {
		return web.ErrNotAcceptable
	}

	hashPassword, err := security.Hash(password)
	if err != nil {
		return err
	}

	dbUser.Password = string(hashPassword)
	dbUser.NeedChangePassword = false

	return s.Repo.Update(ctx, dbUser)
}

// ResetPassword 重置密码
func (s Service) ResetPassword(ctx context.Context, userID uuid.UUID) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	hashPassword, err := security.Hash(DefaultPassword)
	if err != nil {
		return err
	}
	dbUser.Password = string(hashPassword)
	dbUser.NeedChangePassword = true

	return s.Repo.Update(ctx, dbUser)
}

// Lock 锁定
func (s Service) Lock(ctx context.Context, userID uuid.UUID) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	if dbUser.IsLock {
		return web.ErrNotAcceptable
	}

	dbUser.IsLock = true

	return s.Repo.Update(ctx, dbUser)
}

// Unlock 解锁
func (s Service) Unlock(ctx context.Context, userID uuid.UUID) error {
	dbUser, err := s.Repo.FindOne(ctx, userID)
	if err != nil {
		return err
	}

	if !dbUser.IsLock {
		return web.ErrNotAcceptable
	}

	dbUser.IsLock = false

	return s.Repo.Update(ctx, dbUser)
}

// Login 用户登陆
func (s Service) Login(ctx context.Context, userName string, password string) (*User, error) {
	dbUser, err := s.Repo.FindByName(ctx, userName)
	if err != nil {
		return nil, err
	}

	if dbUser.IsLock {
		return nil, web.ErrUserLocked
	}

	err = security.VerifyPassword(dbUser.Password, password)
	if err != nil {
		return nil, web.ErrPasswordIncorrect
	}

	return dbUser, nil
}
