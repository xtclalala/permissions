package system

type M2mRoleMenu struct {
	SysRoleId int `gorm:"column:sys_role_id"`
	SysMenuId int `gorm:"column:sys_menu_id"`
}
