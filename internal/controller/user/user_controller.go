package user

import (
	"bm-novel/internal/controller"
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/auth"
	"bm-novel/internal/infrastructure/persistence"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/joyparty/httpkit"
	"net/http"
)

// 创建用户
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

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	_, err := user.New(userRepo).Create(u)

	if err == nil {
		w.WriteHeader(201)
		return
	}

	controller.WriteStats(w, err)
}

// 查询用户列表
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

	userRepo := persistence.UserRepository{Ctx: r.Context()}

	users, err := userRepo.FindList(params.RoleCode, params.RealName, params.PageIndex, params.PageSize)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	writeUsersResp(users, w)
}

// 编辑用户
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

	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		w.WriteHeader(404)
		return
	}

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := user.New(userRepo).Load(userId)

	if err == nil {
		err = usr.Edit(u)
	}

	controller.WriteStats(w, err)
}

// 锁定用户
func PostUsersLock(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		w.WriteHeader(404)
		return
	}

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := user.New(userRepo).Load(userId)

	if err == nil {
		err = usr.Lock()
	}

	controller.WriteStats(w, err)

}

// 解锁用户
func DeleteUsersLock(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		w.WriteHeader(404)
		return
	}

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := user.New(userRepo).Load(userId)

	if err == nil {
		err = usr.Unlock()
	}

	controller.WriteStats(w, err)
}

// 重置密码
func DeleteUsersPassword(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	if userId == "" {
		w.WriteHeader(404)
		return
	}

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := user.New(userRepo).Load(userId)

	if err == nil {
		err = usr.ResetPassword()
	}

	controller.WriteStats(w, err)
}

// 用户登陆
func PostUsersSession(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 账号名
		UserName string `json:"user_name" valid:"required"`
		// 密码
		Password string `json:"password" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	// 查询用户
	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := userRepo.FindByName(params.UserName)

	if err == nil && usr.IsLock {
		err = user.ErrUserLocked
	}

	if err != nil {
		controller.WriteStats(w, err)
		return
	}

	// 验证密码
	usr.SetRepo(userRepo)
	err = usr.CheckPassword(params.Password)
	if err != nil {
		// 验证失败
		controller.WriteStats(w, err)
		return
	}

	// 下发cookie
	if err = auth.SetAuth(usr, w); err != nil {
		controller.WriteStats(w, err)
		return
	}

	// 下发结构
	writeLoginResp(usr, w)
}

// 登陆后重设密码
func PutUsersSessionPassword(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 密码
		Password string `json:"password" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	userId, err := auth.GetAuth(r)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := user.New(userRepo).Load(userId)

	if err == nil {
		err = usr.ChangeInitPassword(params.Password)
	}

	controller.WriteStats(w, err)
}

// 用户注销
func DeleteUsersSession(w http.ResponseWriter, r *http.Request) {
	auth.ClearAuth(r, w)
}

type userQueryResp struct {
	UserId   string   `json:"user_id"`
	UserName string   `json:"user_name"`
	RoleCode []string `json:"role_code"`
	RealName string   `json:"real_name"`
	Lock     bool     `json:"lock,bool"`
}

func writeLoginResp(usr *user.User, w http.ResponseWriter) {
	rep := &struct {
		UserId             string `json:"user_id"`
		UserName           string `json:"user_name"`
		RealName           string `json:"real_name"`
		NeedChangePassword bool   `json:"need_change_password"`
	}{
		UserId:             usr.UserID.String(),
		UserName:           usr.UserName,
		RealName:           usr.RealName,
		NeedChangePassword: usr.NeedChangePassword,
	}

	b, err := json.Marshal(rep)
	if err != nil {
		controller.WriteStats(w, err)
		return
	}
	_, _ = w.Write(b)
}

func writeUsersResp(users []user.User, w http.ResponseWriter) {

	res := make([]userQueryResp, 0, len(users))
	for _, v := range users {
		re := userQueryResp{
			UserId:   v.UserID.String(),
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
