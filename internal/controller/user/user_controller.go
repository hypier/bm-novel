package user

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/persistence"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joyparty/httpkit"
	"net/http"
)

// 创建用户
func PostUsers(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 用户名
		UserName string `schema:"user_name" valid:"required"`
		// 角色代码
		RoleCode []string `schema:"role_code" valid:"required"`
		// 姓名
		RealName string `schema:"real_name" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	u := user.User{
		UserName: params.UserName,
		RoleCode: params.RoleCode,
		RealName: params.RealName,
	}

	err := u.Construction(&persistence.UserRepository{Ctx: r.Context()}).Create()

	if err == nil {
		w.WriteHeader(201)
		return
	}

	writeStats(w, err)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 角色名
		RoleCode string `schema:"role_code"`
		// 姓名
		RealName string `schema:"real_name"`
		// 当前页码
		PageIndex int `schema:"page_index"`
		// 每页数量
		PageSize int `schema:"page_size"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	userRepo := persistence.UserRepository{Ctx: r.Context()}

	_, err := userRepo.FindList(params.RoleCode, params.RealName, params.PageIndex, params.PageSize)

	if err != nil {
		w.WriteHeader(500)
	}

	return
}

// 编辑用户
func PatchUsers(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 账号名
		UserName string `schema:"user_name" valid:"required"`
		// 角色代码
		RoleCode []string `schema:"role_code" valid:"required"`
		// 姓名
		RealName string `schema:"real_name" valid:"required"`
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

	usr, err := user.New(&persistence.UserRepository{Ctx: r.Context()}, userId)

	if err == nil {
		err = usr.Edit(&u)
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

	usr, err := user.New(&persistence.UserRepository{Ctx: r.Context()}, userId)

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

	usr, err := user.New(&persistence.UserRepository{Ctx: r.Context()}, userId)

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

	usr, err := user.New(&persistence.UserRepository{Ctx: r.Context()}, userId)

	if err == nil {
		err = usr.ResetPassword()
	}

	writeStats(w, err)
}

// 用户登陆
func PostUsersSession(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 账号名
		UserName string `schema:"user_name" valid:"required"`
		// 密码
		Password string `schema:"password" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	userRepo := &persistence.UserRepository{Ctx: r.Context()}
	usr, err := userRepo.FindByName(params.UserName)
	var flag bool

	if err == nil && usr.IsLock {
		err = user.ErrUserLocked
	}

	if err == nil {
		flag, err = usr.CheckPassword(params.Password)
	}

	// todo 下行构建, 下发Cookie, 判断是否已登陆

	fmt.Println(flag, err)

	writeStats(w, err)
}

// 登陆后重设密码
func PutUsersSessionPassword(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 密码
		Password string `schema:"password" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	// todo 获取cookie
	userId := "123"
	if userId == "" {
		w.WriteHeader(404)
		return
	}

	usr, err := user.New(&persistence.UserRepository{Ctx: r.Context()}, userId)

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
}
