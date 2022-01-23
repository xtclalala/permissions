package system

type M2mRoleMenu struct {
	SysRoleId uint `gorm:"column:sys_role_id"`
	SysMenuId uint `gorm:"column:sys_menu_id"`
}
