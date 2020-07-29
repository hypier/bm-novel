package permission

// IPermissionRepository 权限持久化接口
type IPermissionRepository interface {
	FindAll() (Permissions, error)
	Create(permission *Permission) error
	BatchCreate(permissions *Permissions) error
}
