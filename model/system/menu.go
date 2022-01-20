package system

import "permissions/model"

type SysMenu struct {
	model.BaseID
	Pid       uint   `json:"pid" gorm:"comment:父菜单id"`
	Name      string `json:"name" gorm:"comment:路由name;"`
	Path      string `json:"path" gorm:"comment:路由path;"`
	Hidden    bool   `json:"hidden" gorm:"default:true;comment:是否隐藏"`
	Component string `json:"component" gorm:"comment:前端文件路径"`
	Sort      int    `json:"sort" gorm:"default:100;comment:排序"`
	Mate
	SysRoles []*SysRole `json:"roles" gorm:"many2many:m2m_role_menu;"`
	Children []SysMenu  `json:"children" gorm:"-"`
}

type Mate struct {
}
