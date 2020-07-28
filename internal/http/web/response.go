package web

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrNotAcceptable  不接受修改
	ErrNotAcceptable = errors.New("Not Acceptable")
	// ErrServerError 系统错误
	ErrServerError = errors.New("Server Error")

	// ErrUserConflict 用户名重复错误.
	ErrUserConflict = errors.New("User Conflict")
	// ErrUserNotFound 用户不存在.
	ErrUserNotFound = errors.New("User Not Found")
	// ErrUserLocked 用户被锁定.
	ErrUserLocked = errors.New("User Locked")
	// ErrPasswordIncorrect 用户名或密码错误.
	ErrPasswordIncorrect = errors.New("username or password is incorrect")
)

// WriteStats 下行状态码
func WriteStats(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	switch err {
	case ErrUserNotFound:
		w.WriteHeader(404)
	case ErrNotAcceptable:
		w.WriteHeader(406)
	case ErrUserLocked:
		w.WriteHeader(423)
	case ErrUserConflict:
		w.WriteHeader(409)
	case ErrPasswordIncorrect:
		w.WriteHeader(401)
	case ErrServerError:
		w.WriteHeader(500)
	default:
		w.WriteHeader(500)
	}

	fmt.Printf("%+v", err)
}
