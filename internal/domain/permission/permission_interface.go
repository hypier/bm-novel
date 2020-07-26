package permission

type IPermissionRepository interface {
	FindAll() (Permissions, error)
	Create(permission *Permission) error
}
