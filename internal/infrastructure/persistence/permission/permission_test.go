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
	repo := &Repository{ctx}

	all, _ := repo.FindAll()

	fmt.Println(all)
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
	repo := &Repository{ctx}

	err := repo.Create(&per)

	fmt.Println(err)
}
