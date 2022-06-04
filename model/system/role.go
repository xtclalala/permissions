package system

type SysRole struct {
	BaseID
	Name string `json:"name" gorm:"not null"`
	Code string `json:"code" gorm:"not null"`
	Sort int    `json:"sort" gorm:"default:0;not null;"`
	// o2m
	SysOrganizeId int         `json:"organizeId" gorm:"default:0;comment:用户角色ID"`
	SysOrganize   SysOrganize `json:"organize" gorm:"foreignKey:id;references:SysOrganizeId"`
	// m2m
	SysMenus       []SysMenu       `json:"menus" gorm:"many2many:m2m_role_menu;"`
	SysPermissions []SysPermission `json:"permissions" gorm:"many2many:m2m_role_permission;"`
	SysUsers       []SysUser       `json:"users" gorm:"many2many:m2m_user_role;"`
	Children       []SysRole       `json:"children" gorm:"-"`
}
