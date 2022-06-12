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
	var data system.RoleBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.Register(system.SysRole{
		Name: data.Name,
		Code: data.Code,
		Sort: data.Sort,
		SysOrganize: system.SysOrganize{
			BaseID: system.BaseID{
				ID: data.SysOrganizeId,
			},
		},
	}); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}

// UpdateBaseRole 跟新角色基本信息
func (a *RoleApi) UpdateBaseRole(c *gin.Context) {
	var data system.RoleBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.UpdateRoleInfo(system.SysRole{
		BaseID:        system.BaseID{ID: data.Id},
		Name:          data.Name,
		Code:          data.Code,
		Sort:          data.Sort,
		SysOrganizeId: data.SysOrganizeId,
	}); err != nil {
		common.FailWhitStatus(utils.UpdateRoleError, c)
		return
	}
	common.Ok(c)
}

// GetRoleMenus 角色菜单
func (a RoleApi) GetRoleMenus(c *gin.Context) {
	var data system.RoleId
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	err, vos := roleService.GetMenusById(data.Id)
	if err != nil {
		common.Fail(c)
		return
	}
	common.OkWithData(vos.SysMenus, c)
}

// GetRolePer 角色权限
func (a RoleApi) GetRolePer(c *gin.Context) {
	var data system.RoleId
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	err, vos := roleService.GetPermissionsById(data.Id)
	if err != nil {
		common.Fail(c)
		return
	}
	common.OkWithData(vos.SysPermissions, c)
}

// UpdateRoleMenus 更新角色菜单
func (a RoleApi) UpdateRoleMenus(c *gin.Context) {
	var data system.RolePerInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.SetRoleMenu(data.Id, data.SysMenuIds); err != nil {
		common.FailWhitStatus(utils.UpdateRoleMenusError, c)
		return
	}
	common.Ok(c)
}

// UpdateRolePer 更新角色权限
func (a RoleApi) UpdateRolePer(c *gin.Context) {
	var data system.RolePerInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := roleService.SetRolePer(data.Id, data.SysPermissionIds); err != nil {
		common.FailWhitStatus(utils.UpdateRolePerError, c)
		return
	}
	common.Ok(c)
}

// CopyRole 拷贝角色信息
func (a *RoleApi) CopyRole(c *gin.Context) {
	var data system.RoleId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, role := roleService.GetCompleteInfoById(data.Id)
	if err := utils.Validate(&data); err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
		return
	}
	if err = roleService.Register(system.SysRole{
		Name:           role.Name + " Copy",
		Code:           role.Code + "Copy",
		Sort:           role.Sort,
		SysOrganizeId:  role.SysOrganizeId,
		SysOrganize:    role.SysOrganize,
		SysMenus:       role.SysMenus,
		SysPermissions: role.SysPermissions,
	}); err != nil {
		common.FailWhitStatus(utils.CreateRoleError, c)
		return
	}
	common.Ok(c)
}

// CompleteRole 角色详细信息
func (a RoleApi) CompleteRole(c *gin.Context) {
	var data system.RoleId
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	err, role := roleService.GetCompleteInfoById(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
		return
	}
	common.OkWithData(role, c)
}

// DeleteRole 删除角色
func (a *RoleApi) DeleteRole(c *gin.Context) {
	var data system.RoleId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err := roleService.DeleteRole(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeleteRoleError, c)
		return
	}
	common.Ok(c)
}

// SearchRole 搜索角色
func (a *RoleApi) SearchRole(c *gin.Context) {
	var data system.SearchRole
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, list, total := roleService.Search(data)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
		return
	}
	common.OkWithData(&common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

// RoleAllByOrg 查找组织下的所有角色
func (a *RoleApi) RoleAllByOrg(c *gin.Context) {
	data, ok := c.GetQueryArray("ids[]")
	if ok != true {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	err, roles := roleService.GetRoleByOrgIds(data)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
		return
	}
	common.OkWithData(roles, c)
}
