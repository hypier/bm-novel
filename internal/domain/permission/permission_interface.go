package permission

import "context"

// IPermissionRepository 权限持久化接口
type IPermissionRepository interface {
	FindAll(ctx context.Context) (Permissions, error)
	Create(ctx context.Context, permission *Permission) error
	BatchCreate(ctx context.Context, permissions *Permissions) error
}
