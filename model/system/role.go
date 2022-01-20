package system

import "permissions/model"

type SysRole struct {
	model.BaseID
	RoleName string `json:"roleName" gorm:"not null"`
	Pid      uint   `json:"pid" gorm:"default:0"`
	Sort     int    `json:"sort" gorm:"not null;"`
	// 路由
	SysUsers       []*SysUser       `json:"users" gorm:"many2many:m2m_user_role;"`
	SysMenus       []*SysMenu       `json:"menus" gorm:"many2many:m2m_role_menu;"`
	SysPermissions []*SysPermission `json:"permissions" gorm:"many2many:m2m_role_permission;"`
	Children       []SysRole        `json:"children" gorm:"-"`
}
