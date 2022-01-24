package system

import "permissions/model/common"

// SearchUser 分页显示 搜索后的用户
type SearchUser struct {
	common.BasePage
	User
}

// User 创建用户
type User struct {
	UserBaseInfo
	UserPerInfo
}

// UserBaseInfo 用户基本信息
type UserBaseInfo struct {
	Username  string `json:"username"  validate:"max=15,min=6,required" label:"用户名"`
	LoginName string `json:"loginName" validate:"max=15,min=6,required" label:"账号"`
	Password  string `json:"password"  validate:"max=15,min=6,required" label:"密码"`
}

// UserPerInfo 用户权限信息
type UserPerInfo struct {
	SysRoleIds     []uint `json:"roleIds"   validate:"unique,min=1"              label:"角色"`
	SysOrganizeIds []uint `json:"orgIds"    validate:"unique,min=1"              label:"组织"`
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
