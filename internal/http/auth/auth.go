// auth 认证
package auth

import (
	permission2 "bm-novel/internal/domain/permission"
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/cookie"
	"bm-novel/internal/infrastructure/persistence/permission"
	"bm-novel/internal/infrastructure/redis"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	uuid "github.com/satori/go.uuid"
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
	}

	return err
}

func GetAuth(r *http.Request) (userID string, err error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return
	}

	if id, err := claims["id"]; err {
		userID = id.(string)
	}

	return
}

func getAuthUID(r *http.Request) (uid string, err error) {
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
	userID, err := GetAuth(r)
	if err != nil {
		return
	}

	key := fmt.Sprintf("bm:login:%s", userID)

	_ = redis.GetChcher().Delete(key)
	cookie.ClearCookie("jwt", r, w)
}

// LoginAuthenticator 认证
func LoginAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从jwt中获取用户ID
		userID, err := GetAuth(r)
		if err != nil || userID == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// 从redis里获取uid
		key := fmt.Sprintf("bm:login:%s", userID)
		sUID, err := redis.GetChcher().Get(key)
		if err != nil || sUID == nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}
		serverUID := string(sUID)

		// 从jwt中获取用户uid
		clientUID, err := getAuthUID(r)
		if err != nil || clientUID == "" {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// 比较客户端的uid 是否等于 redis中的uid , 不相等则为session过期或被踢掉
		if serverUID != clientUID {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Authorization 授权
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//_ = setRedis(r)

		uri := r.URL.String()
		method := r.Method

		pm, err := getPermission(uri, method)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		fmt.Println(pm)

		next.ServeHTTP(w, r)
	})
}

func getPermission(uri string, method string) (pm *permission2.Permission, err error) {
	key := "bm:permission"
	field := fmt.Sprintf("%s%s", method, uri)

	err = redis.GetChcher().Exists(key, field)
	if err != nil {
		return
	}

	val, err := redis.GetChcher().HGet(key, field)
	if err != nil {
		return
	}

	if val != nil {
		err = json.Unmarshal(val, &pm)
	}

	return
}

func setRedis(r *http.Request) error {
	key := "bm:permission"
	if err := redis.GetChcher().Exists(key); err != nil {
		return nil
	}

	// 设置缓存
	repo := &permission.PermissionRepository{Ctx: r.Context()}
	pms, err := repo.FindAll()
	if err != nil || pms == nil {
		return nil
	}

	for _, v := range *pms {
		ps, err := json.Marshal(v.Users)
		if err != nil {
			continue
		}

		field := fmt.Sprintf("%s%s", v.Method, v.URI)
		err = redis.GetChcher().HPut(key, field, ps, time.Hour*24)
	}

	return err
}
