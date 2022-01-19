package model

type SysMenu struct {
	baseID
	Pid       uint   `json:"pid" gorm:"comment:父菜单id"`
	Name      string `json:"name" gorm:"comment:路由name;not null"`
	Path      string `json:"path" gorm:"comment:路由path;not null"`
	Hidden    bool   `json:"hidden" gorm:"default:true;comment:是否隐藏"`
	Component string `json:"component" gorm:"not null;comment:前端文件路径"`
	Sort      int    `json:"sort" gorm:"default:100;comment:排序"`
	Mate
	SysRoles []SysRole `json:"roles" gorm:"many2many:role_menu;"`
	Children []SysMenu `json:"children" gorm:"-"`
}

type Mate struct {
}
