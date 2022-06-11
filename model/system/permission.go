package system

type SysPermission struct {
	BaseID
	Title string `json:"title" gorm:"not null;"`
	Code  string `json:"code" gorm:"not null;"`
	Sort  int    `json:"sort" gorm:"default:100;comment:排序"`
	// o2m
	SysMenuId int `json:"menuId"`
	// m2m
	SysRoles []SysRole `json:"roles" gorm:"many2many:m2m_role_permission;"`
}
