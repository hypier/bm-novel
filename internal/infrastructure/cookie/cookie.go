package cookie

import (
	"net/http"
)

// 存入cookie,使用cookie存储
func AddCookie(name string, value string, w http.ResponseWriter) {

	cookie := http.Cookie{Name: name, Value: value, Path: "/"}
	http.SetCookie(w, &cookie)
}

// 获取cookie
func GetCookie(name string, r *http.Request) (string, error) {
	cookie, err := r.Cookie(name)

	if err != nil || cookie.Value == "" {
		return "", err
	} else {
		return cookie.Value, nil
	}

}

// 清除cookie
func ClearCookie(name string, r *http.Request, w http.ResponseWriter) bool {
	if _, err := r.Cookie(name); err != nil {
		return false
	}

	cookie := http.Cookie{Name: name, Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)

	return true
}
