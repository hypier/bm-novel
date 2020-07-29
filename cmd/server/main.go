package main

import (
	"bm-novel/internal/config"
	"bm-novel/internal/controller/user"
	"bm-novel/internal/http/auth"
	"fmt"
	"net/http"

	"github.com/joyparty/httpkit"
	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func init() {
	config.LoadConfig()
}

func main() {

	fmt.Printf("Start Server(%s)...\n", config.Config.Server)
	err := http.ListenAndServe(config.Config.Server, APIRouter())

	if err != nil {
		fmt.Println(err)
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
