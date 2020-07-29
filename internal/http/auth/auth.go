package auth

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/http/web"
	"bm-novel/internal/infrastructure/cookie"
	"bm-novel/internal/infrastructure/persistence/permission"
	"bm-novel/internal/infrastructure/redis"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	uuid "github.com/satori/go.uuid"
)

var (
	// TokenAuth jwt认证
	TokenAuth *jwtauth.JWTAuth
	// permissionTime 权限点缓存时间
	permissionTime = time.Hour * 48
	// loginExpHour 登陆过期时间
	loginExpHour = time.Hour * 24
	// visitorCacheKey 访问者的缓存key
	visitorCacheKey = "bm:login:%s"
	// permissionCacheKey 权限缓存key
	permissionCacheKey = "bm:permission"
)

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

// WriteAuth 写入认证信息
func WriteAuth(visitor *user.User, w http.ResponseWriter) error {

	key := fmt.Sprintf(visitorCacheKey, visitor.UserID.String())
	visitID := uuid.NewV4().String()

	logrus.WithFields(logrus.Fields{
		"visitor": visitor,
		"visitID": visitID,
	}).Debug("Write Auth")

	// 可重复写以前数据
	err := redis.GetChcher().Put(key, []byte(visitID), loginExpHour)
	if err != nil {
		return err
	}

	token, err := generateClientToken(visitor, visitID)
	if err != nil {
		return err
	}

	cookie.AddCookie("jwt", token, w)
	return nil
}

// ClearAuth 清除认证
func ClearAuth(r *http.Request, w http.ResponseWriter) error {
	userID, err := GetVisitorUserID(r)
	if err != nil {
		return err
	}

	key := fmt.Sprintf(visitorCacheKey, userID)

	err = redis.GetChcher().Delete(key)
	if err != nil {
		return err
	}

	return cookie.ClearCookie("jwt", r, w)
}

// generateClientToken 生成客户端token
func generateClientToken(visitor *user.User, visitID string) (string, error) {
	claims := jwt.MapClaims{
		"name":  visitor.UserName,
		"id":    visitor.UserID.String(),
		"roles": visitor.RoleCode,
		"jti":   visitID,
		"exp":   time.Now().Add(loginExpHour),
	}
	_, tokenString, err := TokenAuth.Encode(claims)

	logrus.WithFields(logrus.Fields{
		"visitor": visitor,
		"visitID": visitID,
		"token":   tokenString,
	}).Debug("generate Client Token (JWT)", err)

	return tokenString, err
}

// GetVisitorUserID 获取访问者用户ID
func GetVisitorUserID(r *http.Request) (userID uuid.UUID, err error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return
	}

	if id, ok := claims["id"]; ok {
		userID, err = uuid.FromString(id.(string))
	} else {
		err = web.ErrVisitorNotFound
	}

	return
}

// getVisitID 获取访问ID
func getVisitID(r *http.Request) (visitID string, err error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return
	}

	if id, ok := claims["jti"]; ok {
		visitID = id.(string)
	} else {
		err = web.ErrVisitorNotFound
	}

	return
}

// LoginAuthenticator 认证
func LoginAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从jwt中获取用户ID
		userID, err := GetVisitorUserID(r)
		if err != nil {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		// 从redis里获取uid
		key := fmt.Sprintf(visitorCacheKey, userID)
		exists, err := redis.GetChcher().Exists(key)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if !exists {
			http.Error(w, http.StatusText(401), 401)
			return
		}

		sUID, err := redis.GetChcher().Get(key)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		serverUID := string(sUID)

		// 从jwt中获取用户uid
		clientUID, err := getVisitID(r)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		if clientUID == "" {
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
			http.Error(w, http.StatusText(403), 403)
			return
		}

		pattern := chi.RouteContext(r.Context()).RoutePattern()

		if strings.HasSuffix(pattern, "/*") {
			routePath := chi.RouteContext(r.Context()).RoutePath
			pattern = strings.TrimRight(pattern, "/*") + routePath
		}

		fmt.Println(r.URL.String(), pattern)
		method := r.Method

		roles, err := getPermission(pattern, method)
		if err != nil {
			http.Error(w, http.StatusText(403), 403)
			return
		}

		if roles == nil {
			// 没有配置默认通过
			http.Error(w, http.StatusText(403), 403)
			//next.ServeHTTP(w, r)
			return
		}

		// 检查是否有权限
		if !checkRoles(r, roles) {
			http.Error(w, http.StatusText(403), 403)
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
	field := fmt.Sprintf("%s%s", method, uri)

	exists, err := redis.GetChcher().HExists(permissionCacheKey, field)
	if err != nil {
		return
	}

	// 不存在
	if !exists {
		return
	}

	val, err := redis.GetChcher().HGet(permissionCacheKey, field)
	if err != nil {
		return
	}

	if val != nil {
		err = json.Unmarshal(val, &roles)
	}

	return
}

func putCache(r *http.Request) error {

	exists, err := redis.GetChcher().Exists(permissionCacheKey)

	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	repo := permission.New()
	pms, err := repo.FindAll(r.Context())
	if err != nil {
		return err
	}

	num := len(pms)
	if num <= 0 {
		return web.ErrServerError
	}

	var values []interface{}

	for _, v := range pms {
		roles, err := json.Marshal(v.Roles)
		if err != nil {
			continue
		}

		field := fmt.Sprintf("%s%s", v.Method, v.URI)
		values = append(values, field)
		values = append(values, roles)
	}

	err = redis.GetChcher().HMPut(permissionCacheKey, permissionTime, values)

	return err
}
