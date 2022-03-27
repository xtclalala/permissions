package system

type SysRole struct {
	BaseID
	Name string `json:"name" gorm:"not null"`
	Pid  int    `json:"pid" gorm:"default:0"`
	Sort int    `json:"sort" gorm:"not null;"`
	// o2m
	SysOrganizeId int `json:"organizeId"`
	// m2m
	SysMenus       []SysMenu       `json:"menus" gorm:"many2many:m2m_role_menu;"`
	SysPermissions []SysPermission `json:"permissions" gorm:"many2many:m2m_role_permission;"`
	SysUsers       []SysUser       `json:"users" gorm:"many2many:m2m_user_role;"`
	Children       []SysRole       `json:"children" gorm:"-"`
}
