package user

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/http/auth"
	"bm-novel/internal/http/web"
	ur "bm-novel/internal/infrastructure/persistence/user"
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/go-chi/chi"
	"github.com/joyparty/httpkit"
)

var us = user.Service{Repo: ur.New()}

// PostUsers 创建用户
func PostUsers(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 用户名
		UserName string `json:"user_name" valid:"required"`
		// 角色代码
		RoleCode []string `json:"role_code" valid:"required"`
		// 姓名
		RealName string `json:"real_name" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	u := user.User{
		UserName: params.UserName,
		RoleCode: params.RoleCode,
		RealName: params.RealName,
	}

	_, err := us.Create(r.Context(), u)

	if err == nil {
		w.WriteHeader(201)
		return
	}

	web.WriteHttpStats(w, err)
}

// GetUsers 查询用户列表
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := struct {
		// 角色名
		RoleCode []string `json:"role_code"`
		// 姓名
		RealName string `json:"real_name"`
		// 当前页码
		PageIndex int `json:"page_index"`
		// 每页数量
		PageSize int `json:"page_size"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	users, err := ur.New().FindList(r.Context(), params.RoleCode, params.RealName, params.PageIndex, params.PageSize)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	writeUsersResp(users, w)
}

// PatchUsers 编辑用户
func PatchUsers(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 账号名
		UserName string `json:"user_name" valid:"required"`
		// 角色代码
		RoleCode []string `json:"role_code" valid:"required"`
		// 姓名
		RealName string `json:"real_name" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	u := user.User{
		UserName: params.UserName,
		RoleCode: params.RoleCode,
		RealName: params.RealName,
	}

	userID, err := getUserIDForURLParam(r)
	if err != nil {
		web.WriteHttpStats(w, err)
		return
	}

	err = us.Edit(r.Context(), userID, u)
	web.WriteHttpStats(w, err)
}

// PostUsersLock 锁定用户
func PostUsersLock(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDForURLParam(r)
	if err != nil {
		web.WriteHttpStats(w, err)
		return
	}

	err = us.Lock(r.Context(), userID)
	web.WriteHttpStats(w, err)
}

// DeleteUsersLock 解锁用户
func DeleteUsersLock(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDForURLParam(r)
	if err != nil {
		web.WriteHttpStats(w, err)
		return
	}

	err = us.Unlock(r.Context(), userID)
	web.WriteHttpStats(w, err)
}

// DeleteUsersPassword 重置密码
func DeleteUsersPassword(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDForURLParam(r)
	if err != nil {
		web.WriteHttpStats(w, err)
		return
	}

	err = us.ResetPassword(r.Context(), userID)
	web.WriteHttpStats(w, err)
}

// PostUsersSession 用户登陆
func PostUsersSession(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	params := struct {
		// 账号名
		UserName string `json:"user_name" valid:"required"`
		// 密码
		Password string `json:"password" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	// 验证密码
	var usr, err = us.Login(r.Context(), params.UserName, params.Password)
	if err != nil {
		// 验证失败
		web.WriteHttpStats(w, err)
		return
	}

	// 写入认证信息
	if err = auth.WriteAuth(usr, w); err != nil {
		web.WriteHttpStats(w, err)
		return
	}

	// 下发结构
	writeLoginResp(usr, w)
}

// PutUsersSessionPassword 登陆后重设密码
func PutUsersSessionPassword(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 密码
		Password string `json:"password" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	// 此处只取jwt的userID，因为handler已经对jwt与redis的有效性做了判断
	userID, err := auth.GetVisitorUserIDFromJWT(r)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	err = us.ChangeInitPassword(r.Context(), userID, params.Password)

	web.WriteHttpStats(w, err)
}

// DeleteUsersSession 用户注销.
func DeleteUsersSession(w http.ResponseWriter, r *http.Request) {
	if err := auth.ClearAuth(r, w); err != nil {
		web.WriteHttpStats(w, err)
	}
}

type userQueryResp struct {
	UserID   string   `json:"user_id"`
	UserName string   `json:"user_name"`
	RoleCode []string `json:"role_code"`
	RealName string   `json:"real_name"`
	Lock     bool     `json:"lock"`
}

func writeLoginResp(usr *user.User, w http.ResponseWriter) {
	rep := &struct {
		UserID             string `json:"user_id"`
		UserName           string `json:"user_name"`
		RealName           string `json:"real_name"`
		NeedChangePassword bool   `json:"need_change_password"`
	}{
		UserID:             usr.UserID.String(),
		UserName:           usr.UserName,
		RealName:           usr.RealName,
		NeedChangePassword: usr.NeedChangePassword,
	}

	b, err := json.Marshal(rep)
	if err != nil {
		web.WriteHttpStats(w, err)
		return
	}
	_, _ = w.Write(b)
}

func writeUsersResp(users user.Users, w http.ResponseWriter) {

	res := make([]userQueryResp, 0, len(users))
	for _, v := range users {
		re := userQueryResp{
			UserID:   v.UserID.String(),
			UserName: v.UserName,
			RoleCode: v.RoleCode,
			RealName: v.RealName,
			Lock:     v.IsLock,
		}
		res = append(res, re)
	}

	b, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(b)
}

// getUserIDForURLParam 从URL上获取UserId
func getUserIDForURLParam(r *http.Request) (userID uuid.UUID, err error) {
	id := chi.URLParam(r, "user_id")

	if id == "" {
		return userID, web.ErrUserNotFound
	}

	userID, err = uuid.FromString(id)
	if err != nil {
		return userID, web.ErrServerError
	}

	return userID, nil
}
