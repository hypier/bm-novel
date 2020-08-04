package novel

import (
	"fmt"
	"net/http"
)

// GetNovels 获取小说列表
func GetNovels(w http.ResponseWriter, r *http.Request) {
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

	fmt.Println(params)

}
