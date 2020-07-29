package web

import (
	"net/http"

	"github.com/sirupsen/logrus"

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
	// ErrVisitorNotFound 没有找到访问者
	ErrVisitorNotFound = errors.New("Visitor Not Found")
)

// WriteHTTPStats 下行状态码
func WriteHTTPStats(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	switch {
	case errors.Is(err, ErrUserNotFound):
		w.WriteHeader(404)
	case errors.Is(err, ErrNotAcceptable):
		w.WriteHeader(406)
	case errors.Is(err, ErrUserLocked):
		w.WriteHeader(423)
	case errors.Is(err, ErrUserConflict):
		w.WriteHeader(409)
	case errors.Is(err, ErrPasswordIncorrect):
		w.WriteHeader(401)
	case errors.Is(err, ErrServerError):
		w.WriteHeader(500)
	default:
		w.WriteHeader(500)
	}

	panic(err)
	//fmt.Printf("%+v", err)
}

// WriteErrLog 写入错误日志并返回error
func WriteErrLog(err error, format string, args ...interface{}) error {
	logrus.WithError(err).Errorf(format, args...)
	return errors.WithStack(err)
}

// WriteErrLogWithField 写入错误日志并返回error
func WriteErrLogWithField(fields map[string]interface{}, err error, format string, args ...interface{}) error {
	logrus.WithError(err).
		WithFields(fields).
		Errorf(format, args...)

	return errors.WithStack(err)
}
