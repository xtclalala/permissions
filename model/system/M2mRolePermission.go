package system

type M2mRolePermission struct {
	SysRoleId       uint `gorm:"column:sys_role_id"`
	SysPermissionId uint `gorm:"column:sys_permission_id"`
}
