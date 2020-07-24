package main

import (
	"bm-novel/internal/controller/user"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	fmt.Println("Start Server(8888)...")
	_ = http.ListenAndServe(":8888", APIRouter())
}

func APIRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", user.GetUsers)
		r.Post("/", user.PostUsers)

		r.Route("/{user_id}", func(r chi.Router) {
			r.Patch("/", user.PatchUsers)
			r.Post("/lock", user.PostUsersLock)
			r.Delete("/lock", user.DeleteUsersLock)
			r.Delete("/password", user.DeleteUsersPassword)
		})

		r.Route("/session", func(r chi.Router) {
			r.Post("/", user.PostUsersSession)
			r.Delete("/", user.DeleteUsersSession)
			r.Put("/password", user.PutUsersSessionPassword)
		})

	})

	return r
}
