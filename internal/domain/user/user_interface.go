package user

// IUserServer 用户领域服务.
type IUserServer interface {
	Load(userID string) (*User, error)
	Create(user User) (*User, error)
	Edit(user User) error
	ChangeInitPassword(password string) error
	ResetPassword() error
	Lock() error
	Unlock() error
	CheckPassword(password string) error
}

// IUserRepository 用户持久化服务.
type IUserRepository interface {
	FindList(roleCode []string, realName string, pageIndex int, pageSize int) (Users, error)
	FindOne(id string) (*User, error)
	FindByName(name string) (*User, error)
	Create(user *User) error
	Update(user *User) error
}
