package controller

import (
	"bm-novel/internal/infrastructure/persistence"
	"github.com/joyparty/httpkit"
	"net/http"
)

type UserController struct {
	userRepo persistence.UserRepository
}

func (u *UserController) CreatePost(w http.Response, r *http.Request) {
	params := struct {
		// 用户id
		UserId string `schema:"userId" valid:"required"`
		// 用户名
		UserName string `schema:"userName" valid:"required"`
		// 角色代码
		RoleCode string `schema:"roleCode" valid:"required"`
		// 姓名
		TrueName string `schema:"trueName" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

}
