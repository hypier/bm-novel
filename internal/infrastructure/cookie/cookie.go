package cookie

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// AddCookie 存入cookie,使用cookie存储
func AddCookie(name string, value string, w http.ResponseWriter) {

	cookie := http.Cookie{Name: name, Value: value, Path: "/"}
	http.SetCookie(w, &cookie)
}

// GetCookie 获取cookie
func GetCookie(name string, r *http.Request) (string, error) {
	cookie, err := r.Cookie(name)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"name": name,
		}).Debug("cookie clear error ", err)
		return "", err
	}

	return cookie.Value, nil
}

// ClearCookie 清除cookie
func ClearCookie(name string, r *http.Request, w http.ResponseWriter) error {
	if _, err := r.Cookie(name); err != nil {
		logrus.WithFields(logrus.Fields{
			"name": name,
		}).Warn("cookie clear error ", err)
		return err
	}

	cookie := http.Cookie{Name: name, Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)
	return nil
}
