package model

type Permission struct {
	baseID
	Name     string    `json:"name" gorm:"not null;"`
	SysRoles []SysRole `json:"roles" gorm:"many2many:role_permission;"`
}
