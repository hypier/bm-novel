package user

type IUserServer interface {
	Get(userId string) (*User, error)
	Create(user *User) error
	SetRole(roleCode string) error
	InitPassword(password string) error
	ResetPassword() error
	Lock() error
	Unlock() error
	CheckPassword(password string) (bool, error)
}
