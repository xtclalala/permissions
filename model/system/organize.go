package system

type SysOrganize struct {
	BaseID
	Name string `json:"name" gorm:"size:50;not null;comment:组织名称"`
	Code string `json:"code" gorm:"size:50;not null;comment:组织编号"`
	Pid  int    `json:"pid" gorm:"default:0;comment:父级id"`
	Sort int    `json:"sort" gorm:"default:100;comment:排序"`
	// o2m
	SysRoles []SysRole `json:"roles" gorm:"foreignKey:SysOrganizeId"`
	// m2m
	SysUsers []SysUser     `json:"users" gorm:"many2many:m2m_user_organize"`
	Children []SysOrganize `json:"children" gorm:"-"`
}
