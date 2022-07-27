package system

type M2mRoleMenu struct {
	SysRoleId int `gorm:"column:sys_role_id"`
	SysMenuId int `gorm:"column:sys_menu_id"`
}

func (m M2mRoleMenu) TableName() string {
	return "sys_m2m_role_menu"
}
