package user

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/persistence"
	"github.com/joyparty/httpkit"
	"net/http"
)

func PostUsers(w http.ResponseWriter, r *http.Request) {
	params := struct {
		// 用户名
		UserName string `schema:"userName" valid:"required"`
		// 角色代码
		RoleCode []string `schema:"roleCode" valid:"required"`
		// 姓名
		TrueName string `schema:"trueName" valid:"required"`
	}{}

	httpkit.MustScanJSON(&params, r.Body)

	userRepo := persistence.UserRepository{Ctx: r.Context()}
	u := user.User{
		UserName: params.UserName,
		RoleCode: params.RoleCode,
		RealName: params.TrueName,
		Repo:     &userRepo,
	}

	err := u.Create(&u)

	if err == user.ErrUserConflict {
		w.WriteHeader(409)
	} else if err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(201)
	}
}
