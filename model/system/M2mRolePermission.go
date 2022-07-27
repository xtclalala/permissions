package system

type M2mRolePermission struct {
	SysRoleId       int `gorm:"column:sys_role_id"`
	SysPermissionId int `gorm:"column:sys_permission_id"`
}

func (m M2mRolePermission) TableName() string {
	return "sys_m2m_role_permission"
}
