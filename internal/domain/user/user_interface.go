package user

type IUserServer interface {
	Load(userId string) (*User, error)
	Create(user User) (*User, error)
	Edit(user User) error
	ChangeInitPassword(password string) error
	ResetPassword() error
	Lock() error
	Unlock() error
	CheckPassword(password string) error
}

type IUserRepository interface {
	FindList(roleCode []string, realName string, pageIndex int, pageSize int) ([]User, error)
	FindOne(id string) (*User, error)
	FindByName(name string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}
