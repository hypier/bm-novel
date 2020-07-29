package user

import (
	"bm-novel/internal/config"
	"bm-novel/internal/domain/user"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/postgres"
	"context"
	"testing"

	"github.com/sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
)

func init() {
	config.LoadConfigForTest()
	postgres.InitDB()
}

func TestUserRepository_Create(t *testing.T) {
	u1 := uuid.NewV4()
	usr := &user.User{UserID: u1, UserName: "admin", IsLock: false, RoleCode: []string{"admin"}}
	ctx, _ := context.WithCancel(context.Background())
	repo := New()

	_ = repo.Create(ctx, usr)

	dbUser, err := repo.FindOne(ctx, u1)

	t.Log(dbUser, err)

}

func TestUserRepository_FindByName(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := New()

	usr, _ := repo.FindByName(ctx, "chengfa21n")

	t.Log(usr)
}

func TestUserRepository_FindList(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := New()

	list, _ := repo.FindList(ctx, []string{"admin"}, "", 1, 2)

	for _, v := range list {
		t.Log(v.RealName)
	}
}

func TestError(t *testing.T) {

	//logrus.Error(web.ErrPasswordIncorrect)
	//logrus.WithError(web.ErrPasswordIncorrect).Errorf("222222221")
	//logrus.WithFields(logrus.Fields{
	//	"animal": "walrus",
	//}).WithError(web.ErrPasswordIncorrect).Errorf("222222221")

	//err := web.WriteErrLog(web.ErrPasswordIncorrect, "2222222222")
	err := web.WriteErrLogWithField(
		logrus.Fields{
			"animal": "walrus",
		}, web.ErrPasswordIncorrect, "2222222222")

	//err := web.ErrPasswordIncorrect
	//name := "hypier"
	//logrus.WithFields(logrus.Fields{
	//	"name": name,
	//}).Warn("cookie clear error ", err)
	t.Logf("%+v\n", err)
}
