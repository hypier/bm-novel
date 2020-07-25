package auth

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/cookie"
	"bm-novel/internal/infrastructure/redis"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func setJWT(auth *user.User, uid string) (string, error) {
	claims := jwt.MapClaims{"name": auth.UserName,
		"id":  auth.UserID.String(),
		"jti": uid,
		"exp": time.Now().AddDate(0, 0, 1),
	}
	_, tokenString, err := TokenAuth.Encode(claims)
	fmt.Printf("%s\n", tokenString)

	return tokenString, err
}

func SetAuth(auth *user.User, w http.ResponseWriter) error {

	key := fmt.Sprintf("bm:login:%s", auth.UserID.String())
	val := uuid.NewV4().String()

	// 复写以前数据
	err := redis.GetChcher().Put(key, []byte(val), time.Hour*24)
	if err != nil {
		return err
	}

	if token, err := setJWT(auth, val); err == nil {
		cookie.AddCookie("jwt", token, w)
		return nil
	} else {
		return err
	}
}

func GetAuth(r *http.Request) (userId string, err error) {
	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		return
	}

	if id, err := claims["id"]; err {
		userId = id.(string)
	}

	return
}

func getAuthUid(r *http.Request) (uid string, err error) {
	_, claims, err := jwtauth.FromContext(r.Context())

	if err != nil {
		return
	}

	if id, err := claims["jti"]; err {
		uid = id.(string)
	}

	return
}

func ClearAuth(r *http.Request, w http.ResponseWriter) {
	userId, err := GetAuth(r)
	if err != nil {
		return
	}

	key := fmt.Sprintf("bm:login:%s", userId)

	_ = redis.GetChcher().Delete(key)
	cookie.ClearCookie("jwt", r, w)
}

func LoginAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从jwt中获取用户ID
		userId, err := GetAuth(r)
		if err != nil || userId == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// 从redis里获取uid
		key := fmt.Sprintf("bm:login:%s", userId)
		sUid, err := redis.GetChcher().Get(key)
		if err != nil || sUid == nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		serverUid := string(sUid)

		// 从jwt中获取用户uid
		clientUid, err := getAuthUid(r)
		if err != nil || clientUid == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// 比较客户端的uid 是否等于 redis中的uid , 不相等则为session过期或被踢掉
		if serverUid != clientUid {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}