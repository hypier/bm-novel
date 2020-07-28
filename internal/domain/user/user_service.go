package user

import (
	"bm-novel/internal/infrastructure/security"
	"context"
	"github.com/pkg/errors"

	uuid "github.com/satori/go.uuid"
)

var (
	// ErrUserConflict 用户名重复错误.
	ErrUserConflict = errors.New("User Conflict")
	// ErrUserNotFound 用户不存在.
	ErrUserNotFound = errors.New("User Not Found"
	// ErrUserLocked 用户被锁定.
	ErrUserLocked = "User Locked"
	// ErrNotAcceptable 不接受修改.
	ErrNotAcceptable = "Not Acceptable"
	// ErrPasswordIncorrect 用户名或密码错误.
	ErrPasswordIncorrect = "username or password is incorrect"
	// DefaultPassword 默认密码.
	DefaultPassword = "123456"
)

type Service struct {
	repo IUserRepository
}

func (s Service) Create(ctx context.Context, user User) (*User, error) {
	hashPassword, err := security.Hash(DefaultPassword)
	if err != nil {
		return nil, err
	}

	dbUser, err := s.repo.FindByName(ctx, user.UserName)

	if err != nil {
		return nil, err
	} else if dbUser != nil && dbUser.UserName == user.UserName {
		return nil, errors.New(ErrUserConflict)
	}

	u.Password = string(hashPassword)
	u.UserID = uuid.NewV4()
	u.NeedChangePassword = true
	u.UserName = user.UserName
	u.RoleCode = user.RoleCode
	u.RealName = user.RealName
}

func (s Service) Edit(ctx context.Context, user User) error {
	panic("implement me")
}

func (s Service) ChangeInitPassword(ctx context.Context, userID uuid.UUID, password string) error {
	panic("implement me")
}

func (s Service) ResetPassword(ctx context.Context, userID uuid.UUID) error {
	panic("implement me")
}

func (s Service) Lock(ctx context.Context, userID uuid.UUID) error {
	panic("implement me")
}

func (s Service) Unlock(ctx context.Context, userID uuid.UUID) error {
	panic("implement me")
}

func (s Service) CheckPassword(ctx context.Context, userName string, password string) error {
	panic("implement me")
}
