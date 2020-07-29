package user

//
//import (
//	"bm-novel/internal/domain/user"
//	"context"
//	"fmt"
//	"testing"
//
//	uuid "github.com/satori/go.uuid"
//)
//
//func TestUserRepository_Create(t *testing.T) {
//	u1 := uuid.NewV4()
//	usr := &user.User{UserID: u1, UserName: "admin", IsLock: false, RoleCode: []string{"admin"}}
//	ctx, _ := context.WithCancel(context.Background())
//	repo := New()
//
//	_ = repo.Create(ctx, usr)
//
//	dbUser, err := repo.FindOne(ctx, u1)
//
//	fmt.Println(dbUser, err)
//
//}
//
//func TestUserRepository_FindByName(t *testing.T) {
//	ctx, _ := context.WithCancel(context.Background())
//	repo := New()
//
//	usr, _ := repo.FindByName(ctx, "chengfa21n")
//
//	fmt.Println(usr)
//}
//
//func TestUserRepository_FindList(t *testing.T) {
//	ctx, _ := context.WithCancel(context.Background())
//	repo := New()
//
//	list, _ := repo.FindList(ctx, []string{"admin"}, "", 1, 2)
//
//	for _, v := range list {
//		fmt.Println(v.RealName)
//	}
//}
