package system

import (
	"github.com/google/uuid"
	"permissions/model/common"
)

type File struct {
	Id uuid.UUID `form:"id" json:"id" validate:"required"`
}

// SearchUser 分页显示 搜索后的用户
type SearchUser struct {
	common.BasePage
	SysRoleIds     []int  `form:"roleIds[]" json:"roleIds"`
	SysOrganizeIds []int  `form:"orgIds[]" json:"orgIds"`
	Username       string `form:"name" json:"name"`
	LoginName      string `form:"username" json:"username"`
}

// User 创建用户
type User struct {
	UserId
	UserBaseInfo
	UserPerInfo
}

// UserBaseInfo 用户基本信息
type UserBaseInfo struct {
	UserId
	UserLogin
	Username string `json:"username"  validate:"max=15,min=1,required" label:"用户名"`
}

// UserPerInfo 用户权限信息
type UserPerInfo struct {
	UserId
	SysRoleIds     []int `json:"roleIds"   validate:"unique,min=1"          label:"角色"`
	SysOrganizeIds []int `json:"orgIds"    validate:"unique,min=1"          label:"组织"`
}

// UserLogin 用户登录
type UserLogin struct {
	LoginName string `json:"loginName" validate:"max=15,min=5,required" label:"账号"`
	Password  string `json:"password"  validate:"-" label:"密码"`
}

// UserPassword 用户修改密码
type UserPassword struct {
	UserId
	OldPassword      string `json:"oldPassword"  validate:"max=15,min=5,required" label:"旧密码"`
	NewPassword      string `json:"newPassword"  validate:"max=15,min=5,required" label:"新密码"`
	NewPasswordAgain string `json:"NewPasswordAgain"  validate:"max=15,min=5,required" label:"新密码"`
}

// UserId 用户id
type UserId struct {
	Id uuid.UUID `json:"id"        validate:"-"        label:"用户id"`
}

type SearchRole struct {
	common.BasePage
	Name       string `form:"name" json:"name"`
	OrganizeId int    `form:"orgId" json:"orgId"`
}

type Role struct {
	RoleBaseInfo
	RolePerInfo
	RoleId
}

type RoleId struct {
	Id int `form:"id" json:"id" validate:"-" label:"身份主键"`
}

type RoleBaseInfo struct {
	RoleId
	Name          string `form:"name" json:"name"        validate:"max=15,min=6,required" label:"角色名"`
	Code          string `form:"code" json:"code"        validate:"max=15,min=1,required" label:"角色编号"`
	Sort          int    `form:"sort" json:"sort"        validate:"required"              label:"排序"`
	SysOrganizeId int    `form:"orgId" json:"orgId"       validate:"required"              label:"组织"`
}

type RolePerInfo struct {
	RoleId
	SysMenuIds       []int `form:"menuIds"     json:"menuIds"`
	SysPermissionIds []int `form:"permissionIds" json:"permissionIds"`
}

type SearchPermission struct {
	common.BasePage
	Title      string `form:"title" json:"title"      label:"按钮名称"`
	SysMenuId  int    `form:"menuId" json:"menuId"    label:"菜单id"`
	SysRoleIds []int  `form:"sysRoleId" json:"sysRoleId" label:"角色id"`
}

type Permission struct {
	PermissionBaseInfo
	PermissionPerInfo
}

type PermissionId struct {
	Id int `json:"id"  validate:"-" label:"按钮id"`
}

type PermissionBaseInfo struct {
	PermissionId
	Title     string `json:"title"      validate:"required" label:"按钮名称"`
	Code      string `json:"code"      validate:"required" label:"按钮编码"`
	Sort      int    `json:"sort" validate:"required" label:"排序"`
	SysMenuId int    `json:"menuId"    validate:"required" label:"菜单id"`
}

type PermissionPerInfo struct {
	PermissionId
	SysRoleIds []int `json:"sysRoleId" validate:"required" label:"角色id"`
}

type SearchMenu struct {
	common.BasePage
	Path      string `form:"path"`
	Name      string `form:"name"`
	Title     string `form:"title"`
	Component string `form:"component"`
	Hidden    *bool  `form:"hidden"`
}

type Menu struct {
	MenuBaseInfo
	MenuPerInfo
}

type MenuId struct {
	Id int `json:"id" form:"id"  validate:"required" label:"菜单id"`
}

type MenuBaseInfo struct {
	Id        int    `json:"id"  validate:"-" label:"菜单id"`
	Name      string `json:"name"      validate:"required" label:"菜单英文名称"`
	Title     string `json:"title"      validate:"required" label:"菜单中文名称"`
	Path      string `json:"path"      validate:"required" label:"菜单路径"`
	Hidden    *bool  `json:"hidden" validate:"required" label:"是否隐藏"`
	Component string `json:"component" validate:"required" label:"组件地址"`
	Pid       int    `json:"pid"  validate:"-" label:"父级id"`
	Sort      int    `json:"sort" validate:"required" label:"排序"`
	Icon      string `json:"icon" validate:"required" label:"图标"`
}

type MenuPerInfo struct {
	MenuId
	SysRoleIds []int `json:"sysRoleId" validate:"required" label:"角色id"`
}

type SearchOrganize struct {
	common.BasePage
	Name string `form:"name" json:"name"`
}

type Organize struct {
	OrganizeBaseInfo
}

type OrganizeId struct {
	Id int `form:"id" json:"id" validate:"-" label:"组织id"`
}

type OrganizeBaseInfo struct {
	OrganizeId
	Name string `form:"name" json:"name" validate:"required,min=6,max=50" label:"组织名称"`
	Code string `form:"code" json:"code" validate:"required,min=1,max=10" label:"组织编号"`
	Sort int    `form:"sort" json:"sort" validate:"required" label:"排序"`
	Pid  int    `form:"pid"  json:"pid"`
}
