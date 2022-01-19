package model

type SysUser struct {
	baseUUID
	Username  string    `json:"username" gorm:"type:string;size:35;not null"`
	LoginName string    `json:"loginName" gorm:"type:string;size:50;unique;not null"`
	Password  string    `json:"password" gorm:"not null;"`
	SysRoles  []SysRole `json:"roles" gorm:"many2many:user_role;"`
}
