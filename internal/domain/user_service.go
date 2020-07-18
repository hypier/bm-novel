package domain

import "bm-novel/internal/domain/user"

type IUserServer interface {
	Get(userId string)
	Create(user *user.User)
	SetRole(roleCode string)
	InitPassword(password string)
	ResetPassword()
	Lock()
	Unlock()
	CheckPassword(password string)
}
