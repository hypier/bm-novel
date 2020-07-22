package user

type IUserServer interface {
	Create(user *User) error
	SetRole(roleCode []string) error
	InitPassword(password string) error
	ResetPassword() error
	Lock() error
	Unlock() error
	CheckPassword(password string) (bool, error)
}

type IUserRepository interface {
	FindOne(id string) (*User, error)
	FindByName(name string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}
