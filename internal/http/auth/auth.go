package auth

import (
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

var (
	// TokenAuth jwt认证
	TokenAuth *jwtauth.JWTAuth
	// PermissionTime 权限点缓存时间
	PermissionTime = time.Hour * 48
	// LoginExpHour 登陆过期时间
	LoginExpHour = time.Hour * 24
)

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func setJWT(auth *user.User, uid string) (string, error) {
	claims := jwt.MapClaims{"name": auth.UserName,
		"id":    auth.UserID.String(),
		"roles": auth.RoleCode,
		"jti":   uid,
		"exp":   time.Now().Add(LoginExpHour),
	}
	_, tokenString, err := TokenAuth.Encode(claims)
	fmt.Printf("%s\n", tokenString)

	return tokenString, err
}

// SetAuth 写入认证
func SetAuth(auth *user.User, w http.ResponseWriter) error {

	key := fmt.Sprintf("bm:login:%s", auth.UserID.String())
	val := uuid.NewV4().String()

	// 复写以前数据
	err := redis.GetChcher().Put(key, []byte(val), LoginExpHour)
	if err != nil {
		return err
	}

	if token, err := setJWT(auth, val); err == nil {
		cookie.AddCookie("jwt", token, w)
		return nil
	}

	return err
}

// GetAuth 获取认证
func GetAuth(r *http.Request) (userID string, err error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return
	}

	if id, ok := claims["id"]; ok {
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

// ClearAuth 清除认证
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
		if err := putCache(r); err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		uri := r.URL.String()
		method := r.Method

		roles, err := getPermission(uri, method)
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		if roles == nil {
			// 没有配置默认通过
			next.ServeHTTP(w, r)
			return
		}

		// 检查是否有权限
		if !checkRoles(r, roles) {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkRoles(r *http.Request, roles []string) bool {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return false
	}

	rl, ok := claims["roles"]
	if !ok {
		return false
	}

	ris, ok := rl.([]interface{})
	if !ok {
		return false
	}

	for _, v := range ris {
		rs, ok := v.(string)
		if !ok {
			return false
		}

		for _, s := range roles {
			if rs == s {
				return true
			}
		}
	}

	return false
}

func getPermission(uri string, method string) (roles []string, err error) {
	key := "bm:permission"
	field := fmt.Sprintf("%s%s", method, uri)

	exists, err := redis.GetChcher().HExists(key, field)
	if err != nil {
		return
	}

	// 不存在
	if !exists {
		return
	}

	val, err := redis.GetChcher().HGet(key, field)
	if err != nil {
		return
	}

	if val != nil {
		err = json.Unmarshal(val, &roles)
	}

	return
}

func putCache(r *http.Request) error {
	key := "bm:permission"
	if exists, err := redis.GetChcher().Exists(key); err != nil || exists {
		// 报错 或 已缓存
		return err
	}

	// 设置缓存
	repo := &permission.PermissionRepository{Ctx: r.Context()}
	pms, err := repo.FindAll()
	if err != nil || pms == nil {
		return err
	}

	for _, v := range *pms {
		ps, err := json.Marshal(v.Roles)
		if err != nil {
			continue
		}

		field := fmt.Sprintf("%s%s", v.Method, v.URI)
		err = redis.GetChcher().HPut(key, field, ps, PermissionTime)
	}

	return err
}
