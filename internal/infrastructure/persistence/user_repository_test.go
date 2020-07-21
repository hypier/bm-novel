package persistence

import (
	"bm-novel/internal/domain/user"
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	u1 := uuid.NewV4()
	usr := &user.User{UserId: u1.String(), UserName: "2222", UserLock: true}
	ctx, _ := context.WithCancel(context.Background())
	repo := &UserRepository{ctx: ctx}

	_ = repo.Create(usr)

	dbUser, err := repo.FindOne(u1.String())

	fmt.Println(dbUser, err)

}

func TestUserRepository_FindByName(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := &UserRepository{ctx: ctx}

	usr, _ := repo.FindByName("123")

	fmt.Println(usr)
}
