package permission

import (
	"bm-novel/internal/domain/permission"
	"context"
	"fmt"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestPermissionRepository_FindAll(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	repo := &PermissionRepository{ctx}

	all, _ := repo.FindAll()

	fmt.Println(all)
}

func TestPermissionRepository_Create(t *testing.T) {
	per := permission.Permission{
		PID:    uuid.NewV4(),
		URI:    "/users",
		Name:   "创建",
		Method: "POST",
		Users:  []string{"admin", "common"},
	}

	ctx, _ := context.WithCancel(context.Background())
	repo := &PermissionRepository{ctx}

	_ = repo.Create(&per)
}
