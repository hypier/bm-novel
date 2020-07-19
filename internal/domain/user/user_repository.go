package user

type IUserRepository interface {
	FindOne(id string) (user *User, err error)
	FindByName(name string) (user *User, err error)
	Create(user *User) error
	Update(user *User) error
}
