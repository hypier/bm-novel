package user

import (
	"bm-novel/internal/domain/user"
	"context"
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUserRepository_Create(t *testing.T) {
	u1 := uuid.NewV4()
	usr := &user.User{UserID: u1, UserName: "2222", IsLock: true}
	ctx, _ := context.WithCancel(context.Background())
	repo := &UserRepository{Ctx: ctx}

	_ = repo.Create(usr)

	dbUser, err := repo.FindOne(u1.String())

	fmt.Println(dbUser, err)

}

func TestUserRepository_FindByName(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := &UserRepository{Ctx: ctx}

	usr, _ := repo.FindByName("chengfa21n")

	fmt.Println(usr)
}

func TestUserRepository_FindList(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := &UserRepository{Ctx: ctx}

	list, _ := repo.FindList([]string{"admin"}, "", 1, 2)

	for _, v := range list {
		fmt.Println(v.RealName)
	}
}
