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
	var data system2.PermissionBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := permissionService.Register(&system2.SysPermission{
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils2.CreatePermissionError, c)
		return
	}
	common.Ok(c)
}

// UpdatePerBaseInfo 更新按钮基本信息
func (a *PermissionApi) UpdatePerBaseInfo(c *gin.Context) {
	var data system2.PermissionBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := permissionService.Update(system2.SysPermission{
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils2.UpdatePermissionError, c)
		return
	}
	common.Ok(c)
}

// PermissionAllByMenuId 返回页面所有按钮
func (a *PermissionApi) PermissionAllByMenuId(c *gin.Context) {
	var data system2.MenuId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err, pers := permissionService.GetPerByMenuId(data.Id)
	if err != nil {
		common.FailWhitStatus(utils2.FindPermissionError, c)
		return
	}
	common.OkWithData(pers, c)
}

// DeletePermission 删除按钮
func (a *PermissionApi) DeletePermission(c *gin.Context) {
	var data system2.PermissionId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err := permissionService.DeletePermission(data.Id)
	if err != nil {
		common.FailWhitStatus(utils2.DeletePermissionError, c)
		return
	}
	common.Ok(c)
}
