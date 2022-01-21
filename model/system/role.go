package system

import "permissions/model"

type SysRole struct {
	model.BaseID
	Name           string          `json:"name" gorm:"not null"`
	Pid            uint            `json:"pid" gorm:"default:0"`
	Sort           int             `json:"sort" gorm:"not null;"`
	SysUsers       []SysUser       `json:"users" gorm:"many2many:m2m_user_role;"`
	SysMenus       []SysMenu       `json:"menus" gorm:"many2many:m2m_role_menu;"`
	SysPermissions []SysPermission `json:"permissions" gorm:"many2many:m2m_role_permission;"`
	SysOrganizes   []SysOrganize   `json:"organizes" gorm:"many2many:m2m_organize_role"`
	Children       []SysRole       `json:"children" gorm:"-"`
}
