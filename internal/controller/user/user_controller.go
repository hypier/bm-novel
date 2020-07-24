package user

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/persistence"
	"encoding/json"
	"fmt"
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

	writeStats(w, err)
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
	resp := toUserQueryResp(users)
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	_, _ = w.Write(b)
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

	writeStats(w, err)
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

	writeStats(w, err)

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

	writeStats(w, err)
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

	writeStats(w, err)
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
	var flag bool
	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := userRepo.FindByName(params.UserName)

	if err == nil && usr.IsLock {
		err = user.ErrUserLocked
	} else if err != nil {
		writeStats(w, err)
		return
	}

	// 验证密码
	usr.SetRepo(userRepo)
	flag, err = usr.CheckPassword(params.Password)
	// todo 下行构建, 下发Cookie, 判断是否已登陆

	fmt.Println(flag, err)

}

// 登陆后重设密码
func PutUsersSessionPassword(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 密码
		Password string `json:"password" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	// todo 获取cookie
	userId := "123"
	if userId == "" {
		w.WriteHeader(404)
		return
	}

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := user.New(userRepo).Load(userId)

	if err == nil {
		err = usr.ChangeInitPassword(params.Password)
	}

	writeStats(w, err)
}

// 用户注销
func DeleteUsersSession(w http.ResponseWriter, r *http.Request) {

	// todo 获取cookie，删除cookie
}

func writeStats(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	switch err {
	case user.ErrUserNotFound:
		w.WriteHeader(404)
	case user.ErrNotAcceptable:
		w.WriteHeader(406)
	case user.ErrUserLocked:
		w.WriteHeader(423)
	case user.ErrUserConflict:
		w.WriteHeader(409)
	default:
		w.WriteHeader(500)
	}

	fmt.Println(err)
}

type userQueryResp struct {
	UserId   string   `json:"user_id"`
	UserName string   `json:"user_name"`
	RoleCode []string `json:"role_code"`
	RealName string   `json:"real_name"`
	Lock     bool     `json:"lock,bool"`
}

func toUserQueryResp(userList []user.User) []userQueryResp {

	res := make([]userQueryResp, 0, len(userList))
	for _, v := range userList {
		re := userQueryResp{
			UserId:   v.UserID.String(),
			UserName: v.UserName,
			RoleCode: v.RoleCode,
			RealName: v.RealName,
			Lock:     v.IsLock,
		}
		res = append(res, re)
	}
	return res
}
