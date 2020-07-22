package user

import (
	"bm-novel/internal/domain/user"
	"bm-novel/internal/infrastructure/persistence"
	"fmt"
	"github.com/joyparty/httpkit"
	"net/http"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
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

	userRepo := persistence.UserRepository{Ctx: r.Context()}
	u := user.User{UserId: params.UserId,
		UserName: params.UserName,
		RoleCode: params.RoleCode,
		RealName: params.TrueName,
		Repo:     &userRepo,
	}

	err := u.Create(&u)

	if err != nil {
		fmt.Println(err)
	}
}
