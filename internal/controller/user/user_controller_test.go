package user

import (
	"bm-novel/internal/config"
	"bm-novel/internal/infrastructure/postgres"
	"bm-novel/internal/infrastructure/redis"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

var r *chi.Mux

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {

	logrus.SetLevel(logrus.FatalLevel)
	config.LoadConfigForTest()
	postgres.InitDB()
	redis.InitRedis()

	r = chi.NewMux()
	r.Post("/users/session", PostUsersSession)
}

type userSession struct {
	// 账号名
	UserName string
	// 密码
	Password string
}

func TestPostUsersSession(t *testing.T) {
	params := []struct {
		param  userSession
		status int
	}{
		{
			param:  userSession{UserName: "heyong", Password: "1232456"},
			status: 401,
		},
		{
			param:  userSession{UserName: "heyong", Password: "123456"},
			status: 200,
		},
	}

	for _, table := range params {

		user, _ := json.Marshal(table.param)
		userJSON := strings.NewReader(string(user))

		request, err := http.NewRequest("POST", "/users/session", userJSON)
		if err != nil {
			t.Error(err)
		}

		writer := httptest.NewRecorder()
		r.ServeHTTP(writer, request)
		if writer.Code != table.status {
			t.Error(writer.Body)
		}
	}
}
