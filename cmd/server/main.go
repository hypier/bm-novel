package main

import (
	"bm-novel/internal/config"
	"bm-novel/internal/controller/user"
	"bm-novel/internal/http/auth"
	"bm-novel/internal/infrastructure/postgres"
	"bm-novel/internal/infrastructure/redis"
	"flag"
	"net/http"

	"github.com/joyparty/httpkit"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func init() {
	// 配置文件
	ConfFile := flag.String("config", "./configs/server/config.toml", "config file")
	flag.Parse()

	config.LoadConfig(*ConfFile)

	if level, err := logrus.ParseLevel(config.Config.LogLevel); err == nil {
		logrus.SetLevel(level)
	}

	logrus.Debugf("config file: %s", *ConfFile)

	postgres.InitDB()
	redis.InitRedis()
}

func main() {

	logrus.Infof("Start Server(%s)...", config.Config.Server)

	err := http.ListenAndServe(config.Config.Server, APIRouter())

	if err != nil {
		logrus.Fatal("ListenAndServe, ", err)
	}
}

// APIRouter Api路由
func APIRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(httpkit.Recoverer(logrus.New()))

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(auth.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(auth.LoginAuthenticator)

		r.Route("/users", func(r chi.Router) {
			r.Use(auth.Authorization)

			r.Get("/", user.GetUsers)
			r.Post("/", user.PostUsers)
		})

		r.Route("/users/session", func(r chi.Router) {
			r.Delete("/", user.DeleteUsersSession)
			r.Put("/password", user.PutUsersSessionPassword)
		})

		r.Route("/users/{user_id}", func(r chi.Router) {
			r.Use(auth.Authorization)

			r.Patch("/", user.PatchUsers)
			r.Post("/lock", user.PostUsersLock)
			r.Delete("/lock", user.DeleteUsersLock)
			r.Delete("/password", user.DeleteUsersPassword)
		})

	})

	r.Post("/users/session", user.PostUsersSession)

	return r
}
