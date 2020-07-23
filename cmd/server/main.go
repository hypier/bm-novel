package main

import (
	"bm-novel/internal/controller/user"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Post("/user", user.PostUsers)

	fmt.Println("Start Server(8888)...")
	http.ListenAndServe(":8888", r)
}
