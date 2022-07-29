package system

type SysUser struct {
	BaseUUID
	Username  string `json:"username"  gorm:"type:string;size:35;not null"`
	LoginName string `json:"loginName" gorm:"type:string;size:50;unique;not null"`
	Password  string `json:"password"  gorm:"not null;"`
	// m2m
	SysRoles     []SysRole     `json:"roles"      gorm:"many2many:sys_m2m_user_role;"`
	SysOrganizes []SysOrganize `json:"organizes"  gorm:"many2many:sys_m2m_user_organize"`
}
