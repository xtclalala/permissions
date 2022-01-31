package system

import (
	"github.com/google/uuid"
	"permissions/model/common"
)

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
	UserId
	UserLogin
	Username string `json:"username"  validate:"max=15,min=6,required" label:"用户名"`
}

// UserPerInfo 用户权限信息
type UserPerInfo struct {
	UserId
	SysRoleIds     []uint `json:"roleIds"   validate:"unique,min=1"          label:"角色"`
	SysOrganizeIds []uint `json:"orgIds"    validate:"unique,min=1"          label:"组织"`
}

// UserLogin 用户登录
type UserLogin struct {
	LoginName string `json:"loginName" validate:"max=15,min=6,required" label:"账号"`
	Password  string `json:"password"  validate:"max=15,min=6,required" label:"密码"`
}

// UserId 用户id
type UserId struct {
	Id uuid.UUID `json:"id"        validate:"-"        label:"用户id"`
}

type SearchRole struct {
	common.BasePage
	Role
}

type Role struct {
	RoleBaseInfo
	RolePerInfo
	RoleId
}

type RoleId struct {
	Id uint `json:"id" validate:"-" label:"身份主键"`
}

type RoleBaseInfo struct {
	RoleId
	Name          string `json:"name"        validate:"max=15,min=6,required" label:"角色名"`
	Sort          int    `json:"sort"        validate:"required"              label:"排序"`
	SysOrganizeId uint   `json:"orgId"       validate:"required"              label:"组织"`
	Pid           uint   `json:"pid"`
}

type RolePerInfo struct {
	RoleId
	SysMenuIds       []uint `json:"menuIds"`
	SysPermissionIds []uint `json:"permissions"`
}

type SearchPermission struct {
	common.BasePage
	Permission
}

type Permission struct {
	PermissionBaseInfo
	PermissionPerInfo
}

type PermissionId struct {
	Id uint `json:"id"  validate:"-" label:"按钮id"`
}

type PermissionBaseInfo struct {
	PermissionId
	Name      string `json:"name"      validate:"required" label:"按钮名称"`
	Sort      int    `json:"sort" validate:"required" label:"排序"`
	SysMenuId uint   `json:"menuId"    validate:"required" label:"菜单id"`
}

type PermissionPerInfo struct {
	PermissionId
	SysRoleIds []uint `json:"sysRoleId" validate:"required" label:"角色id"`
}

type SearchMenu struct {
	common.BasePage
	Menu
}

type Menu struct {
	MenuBaseInfo
	MenuPerInfo
}

type MenuId struct {
	Id uint `json:"id"  validate:"-" label:"菜单id"`
}

type MenuBaseInfo struct {
	MenuId
	Name      string `json:"name"      validate:"required" label:"菜单名称"`
	Path      string `json:"path"      validate:"required" label:"菜单路径"`
	Hidden    bool   `json:"hidden" validate:"required" label:"是否隐藏"`
	Component string `json:"component" validate:"required" label:"组件地址"`
	Pid       uint   `json:"pid"  validate:"required" label:"父级id"`
	Sort      int    `json:"sort" validate:"required" label:"排序"`
	Icon      string `json:"icon" validate:"required" label:"图标"`
}

type MenuPerInfo struct {
	MenuId
	SysRoleIds []uint `json:"sysRoleId" validate:"required" label:"角色id"`
}

type SearchOrganize struct {
	common.BasePage
	Organize
}

type Organize struct {
	OrganizeBaseInfo
}

type OrganizeId struct {
	Id uint `json:"id" validate:"-" label:"组织id"`
}

type OrganizeBaseInfo struct {
	OrganizeId
	Name string `json:"name" validate:"required,min:6,max:50" label:"组织名称"`
	Sort int    `json:"sort" validate:"required" label:"排序"`
	Pid  uint   `json:"pid" `
}
