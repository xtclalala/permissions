package system

import "permissions/model/common"

type SearchUser struct {
	common.BasePage
	User
}

type User struct {
	Username  string `json:"username"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

type SearchRole struct {
	common.BasePage
	Role
}

type Role struct {
	Name string `json:"name"`
	Pid  uint   `json:"pid"`
}

type SearchPermission struct {
	common.BasePage
	Permission
}

type Permission struct {
	Name      string `json:"name"`
	SysMenuId uint   `json:"menuId"`
}

type SearchOrganize struct {
	common.BasePage
	Organize
}

type Organize struct {
	Name string `json:"name" gorm:"size:50;not null"`
	Pid  uint   `json:"pid" gorm:"default:0"`
}
