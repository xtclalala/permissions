package system

type SysMenu struct {
	BaseID
	Name      string `json:"name" gorm:"comment:路由name;"`
	Title     string `json:"title" gorm:"comment:路由标题;"`
	Path      string `json:"path" gorm:"comment:路由path;"`
	Hidden    bool   `json:"hidden" gorm:"default:false;comment:是否隐藏"`
	Component string `json:"component" gorm:"comment:前端文件路径"`
	Pid       int    `json:"pid" gorm:"comment:父菜单id"`
	Sort      int    `json:"sort" gorm:"default:100;comment:排序"`
	Mate
	// o2m
	SysPermissions []SysPermission `json:"permissions" gorm:"foreignKey:SysMenuId"`
	// m2m
	SysRoles []SysRole `json:"roles" gorm:"many2many:m2m_role_menu;"`
	Children []SysMenu `json:"children" gorm:"-"`
}

type Mate struct {
	Icon string `json:"icon" gorm:"comment:图标"`
}
