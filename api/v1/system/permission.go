package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
)

type PermissionApi struct{}

// Register 注册页面按钮
func (a *PermissionApi) Register(c *gin.Context) {
	var data system.Permission
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := permissionService.Register(&system.SysPermission{
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils.CreatePermissionError, c)
	}
	common.Ok(c)
}

// UpdatePerBaseInfo 更新按钮基本信息
func (a *PermissionApi) UpdatePerBaseInfo(c *gin.Context) {
	var data system.PermissionBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := permissionService.Update(system.SysPermission{
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils.UpdatePermissionError, c)
	}
	common.Ok(c)
}

// PermissionAllByMenuId 返回页面所有按钮
func (a *PermissionApi) PermissionAllByMenuId(c *gin.Context) {
	var data system.MenuId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, pers := permissionService.GetPerByMenuId(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindPermissionError, c)
	}
	common.OkWithData(pers, c)
}

// DeletePermission 删除按钮
func (a *PermissionApi) DeletePermission(c *gin.Context) {
	var data system.PermissionId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err := permissionService.DeletePermission(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeletePermissionError, c)
	}
	common.Ok(c)
}
