package user

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

// IUserService 用户领域服务.
type IUserService interface {
	Create(ctx context.Context, user *User) error
	Edit(ctx context.Context, userID uuid.UUID, user User) error

	ChangeInitPassword(ctx context.Context, userID uuid.UUID, password string) error
	ResetPassword(ctx context.Context, userID uuid.UUID) error
	Lock(ctx context.Context, userID uuid.UUID) error
	Unlock(ctx context.Context, userID uuid.UUID) error
	Login(ctx context.Context, userName string, password string) (*User, error)
}

// IUserRepository 用户持久化服务.
type IUserRepository interface {
	FindList(ctx context.Context, roleCode []string, realName string, pageIndex int, pageSize int) (Users, error)
	FindOne(ctx context.Context, userID uuid.UUID) (*User, error)
	FindByName(ctx context.Context, name string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
}
