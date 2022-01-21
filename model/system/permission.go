package system

import "permissions/model"

type SysPermission struct {
	model.BaseID
	Name      string    `json:"name" gorm:"not null;"`
	Sort      int       `json:"sort" gorm:"default:100;comment:排序"`
	SysMenuId uint      `json:"menuId"`
	SysRoles  []SysRole `json:"roles" gorm:"many2many:m2m_role_permission;"`
}
