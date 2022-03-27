package system

type M2mRolePermission struct {
	SysRoleId       int `gorm:"column:sys_role_id"`
	SysPermissionId int `gorm:"column:sys_permission_id"`
}
