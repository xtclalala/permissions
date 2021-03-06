package system

import (
	"permissions/services"
)

type SysApiGroup struct {
	UserApi
	RoleApi
	OrganizeApi
	PermissionApi
	MenuApi
	FileApi
}

var menuService = services.ServiceGroupApp.SysServiceGroup.MenuService
var userService = services.ServiceGroupApp.SysServiceGroup.UserService
var roleService = services.ServiceGroupApp.SysServiceGroup.RoleService
var permissionService = services.ServiceGroupApp.SysServiceGroup.PermissionService
var organizeService = services.ServiceGroupApp.SysServiceGroup.OrganizeService
var fileService = services.ServiceGroupApp.SysServiceGroup.FileService
