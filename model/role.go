package model

type SysRole struct {
	baseID
	RoleName string `json:"roleName" gorm:"not null"`
	Pid      uint   `json:"pid" gorm:"default:0"`
	Sort     int    `json:"sort" gorm:"not null;"`
	// 路由
	SysUsers []SysUser `json:"users" gorm:"many2many:user_role;"`
	//SysMenus    []SysMenu    	`json:"menus" gorm:"many2many:role_menu;"`
	//Permissions []Permission 	`json:"permissions" gorm:"many2many:role_permission;"`
	Children []SysRole `json:"children" gorm:"-"`
}
