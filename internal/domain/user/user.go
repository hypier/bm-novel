package user

// 用户基本信息
type User struct {
	// 用户id
	UserId string `json:"userId"`
	// 用户名
	UserName string `json:"userName"`
	// 是否锁定
	Lock bool `json:"lock"`
	// 密码
	Password string `json:"password"`
	// 盐值
	Salt string `json:"sale"`
	// 角色代码
	RoleCode string `json:"roleCode"`
	// 姓名
	TrueName string `json:"trueName"`
}
