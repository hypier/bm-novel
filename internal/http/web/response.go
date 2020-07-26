package web

import (
	"bm-novel/internal/domain/user"
	"fmt"
	"net/http"
)

func WriteStats(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	switch err {
	case user.ErrUserNotFound:
		w.WriteHeader(404)
	case user.ErrNotAcceptable:
		w.WriteHeader(406)
	case user.ErrUserLocked:
		w.WriteHeader(423)
	case user.ErrUserConflict:
		w.WriteHeader(409)
	case user.ErrPasswordIncorrect:
		w.WriteHeader(401)
	default:
		w.WriteHeader(500)
	}

	fmt.Println(err)
}
