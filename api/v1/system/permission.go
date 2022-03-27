package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	system2 "permissions/model/system"
	utils2 "permissions/utils"
)

type PermissionApi struct{}

// Register 注册页面按钮
func (a *PermissionApi) Register(c *gin.Context) {
	var data system2.Permission
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := permissionService.Register(&system2.SysPermission{
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils2.CreatePermissionError, c)
	}
	common.Ok(c)
}

// UpdatePerBaseInfo 更新按钮基本信息
func (a *PermissionApi) UpdatePerBaseInfo(c *gin.Context) {
	var data system2.PermissionBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := permissionService.Update(system2.SysPermission{
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils2.UpdatePermissionError, c)
	}
	common.Ok(c)
}

// PermissionAllByMenuId 返回页面所有按钮
func (a *PermissionApi) PermissionAllByMenuId(c *gin.Context) {
	var data system2.MenuId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, pers := permissionService.GetPerByMenuId(data.Id)
	if err != nil {
		common.FailWhitStatus(utils2.FindPermissionError, c)
	}
	common.OkWithData(pers, c)
}

// DeletePermission 删除按钮
func (a *PermissionApi) DeletePermission(c *gin.Context) {
	var data system2.PermissionId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err := permissionService.DeletePermission(data.Id)
	if err != nil {
		common.FailWhitStatus(utils2.DeletePermissionError, c)
	}
	common.Ok(c)
}
