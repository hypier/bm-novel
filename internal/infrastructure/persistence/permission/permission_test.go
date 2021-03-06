package permission

import (
	"bm-novel/internal/config"
	"bm-novel/internal/domain/permission"
	"bm-novel/internal/infrastructure/postgres"
	"context"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func init() {
	config.LoadConfigForTest()
	postgres.InitDB()
}

func TestPermissionRepository_FindAll(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := New()

	all, _ := repo.FindAll(ctx)

	t.Log(all)
}

func TestPermissionRepository_Create(t *testing.T) {
	per := permission.Permission{
		PID:    uuid.NewV4(),
		URI:    "/users",
		Name:   "查询",
		Method: "GET",
		Roles:  []string{"admin"},
	}

	ctx, _ := context.WithCancel(context.Background())
	repo := New()

	err := repo.Create(ctx, &per)

	t.Log(err)
}

func TestRepository_BatchCreate(t *testing.T) {
	//outsourcer := "Outsourcer"
	responsibleEditor := "ResponsibleEditor"
	chiefEditor := "ChiefEditor"
	admin := "Admin"

	pers := permission.Permissions{
		{PID: uuid.NewV4(), Name: "用户查询列表", URI: "/users", Method: "GET", Roles: []string{admin, chiefEditor, responsibleEditor}},
		{PID: uuid.NewV4(), Name: "创建用户", URI: "/users", Method: "POST", Roles: []string{admin, chiefEditor}},
		{PID: uuid.NewV4(), Name: "编辑用户", URI: "/users/{user_id}", Method: "PATCH", Roles: []string{admin, chiefEditor}},
		{PID: uuid.NewV4(), Name: "用户锁定", URI: "/users/{user_id}/lock", Method: "POST", Roles: []string{admin, chiefEditor}},
		{PID: uuid.NewV4(), Name: "用户解锁", URI: "/users/{user_id}/lock", Method: "DELETE", Roles: []string{admin, chiefEditor}},
		{PID: uuid.NewV4(), Name: "重置密码", URI: "/users/{user_id}/password", Method: "DELETE", Roles: []string{admin, chiefEditor}},
		//{PID: uuid.NewV4(), Name: "用户注销", URI: "/users/session", Method: "DELETE", Roles: []string{}},
		//{PID: uuid.NewV4(), Name: "登陆后重设密码", URI: "/users/session/password", Method: "PUT", Roles: []string{}},
	}

	ctx, _ := context.WithCancel(context.Background())
	repo := New()
	_ = repo.BatchCreate(ctx, &pers)
}
