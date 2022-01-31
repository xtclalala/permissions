package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
)

type RoleApi struct{}

// CreateRole 创建角色
func (a *RoleApi) CreateRole(c *gin.Context) {
	var data system.Role
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	menus := menuService.Ids2Object(data.SysMenuIds)
	pers := permissionService.Ids2Object(data.SysPermissionIds)
	if err := roleService.Register(system.SysRole{
		Name:           data.Name,
		Pid:            data.Pid,
		Sort:           data.Sort,
		SysOrganizeId:  data.SysOrganizeId,
		SysMenus:       menus,
		SysPermissions: pers,
	}); err != nil {
		common.FailWithMessage(err.Error(), c)
	}
	common.Ok(c)
}

// UpdateBaseRole 跟新角色基本信息
func (a *RoleApi) UpdateBaseRole(c *gin.Context) {
	var data system.RoleBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code != utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := roleService.UpdateRoleInfo(system.SysRole{
		BaseID: system.BaseID{ID: data.Id},
		Name:   data.Name,
		Pid:    data.Pid,
		Sort:   data.Sort,
	}); err != nil {
		common.FailWhitStatus(utils.UpdateRoleError, c)
	}
	common.Ok(c)
}

// UpdateRoleMenus 更新角色菜单
func (a RoleApi) UpdateRoleMenus(c *gin.Context) {
	var data system.RolePerInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code != utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := roleService.SetRoleMenu(data.Id, data.SysMenuIds); err != nil {
		common.FailWhitStatus(utils.UpdateRoleMenusError, c)
	}
	if err := roleService.SetRolePer(data.Id, data.SysPermissionIds); err != nil {
		common.FailWhitStatus(utils.UpdateRolePerError, c)
	}
	common.Ok(c)
}

// CopyRole 拷贝角色信息
func (a *RoleApi) CopyRole(c *gin.Context) {
	var data system.Role
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code != utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, role := roleService.GetCompleteInfoById(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
	}
	if err = roleService.Register(system.SysRole{
		Name:           data.Name,
		Pid:            data.Pid,
		Sort:           data.Sort,
		SysOrganizeId:  data.SysOrganizeId,
		SysMenus:       role.SysMenus,
		SysPermissions: role.SysPermissions,
	}); err != nil {
		common.FailWhitStatus(utils.CreateRoleError, c)
	}
	common.Ok(c)
}

func (a RoleApi) CompleteRole(c *gin.Context) {
	var data system.RoleId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code != utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, role := roleService.GetCompleteInfoById(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
	}
	common.OkWithData(role, c)
}

// DeleteRole 删除角色
func (a *RoleApi) DeleteRole(c *gin.Context) {
	var data system.RoleId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code != utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err := roleService.DeleteRole(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeleteRoleError, c)
	}
	common.Ok(c)
}

// SearchRole 搜索角色
func (a *RoleApi) SearchRole(c *gin.Context) {
	var data system.SearchRole
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, list, total := roleService.Search(data)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
	}
	common.OkWithData(&common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

// RoleAllByOrg 查找组织下的所有用户
func (a *RoleApi) RoleAllByOrg(c *gin.Context) {
	var data system.OrganizeId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, roles := roleService.GetRoleByOrgId(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
	}
	common.OkWithData(roles, c)
}
