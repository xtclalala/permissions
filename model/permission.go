package model

type SysPermission struct {
	BaseID
	Name     string     `json:"name" gorm:"not null;"`
	SysRoles []*SysRole `json:"roles" gorm:"many2many:m2m_role_permission;"`
}
