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
	var data system.PermissionBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := permissionService.Register(&system.SysPermission{
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils.CreatePermissionError, c)
		return
	}
	common.Ok(c)
}

// UpdatePerBaseInfo 更新按钮基本信息
func (a *PermissionApi) UpdatePerBaseInfo(c *gin.Context) {
	var data system.PermissionBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := permissionService.Update(system.SysPermission{
		BaseID: system.BaseID{
			ID: data.Id,
		},
		Name:      data.Name,
		Sort:      data.Sort,
		SysMenuId: data.SysMenuId,
	}); err != nil {
		common.FailWhitStatus(utils.UpdatePermissionError, c)
		return
	}
	common.Ok(c)
}

// PermissionAllByMenuId 返回页面所有按钮
func (a *PermissionApi) PermissionAllByMenuId(c *gin.Context) {
	var data system.MenuId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err, pers := permissionService.GetPerByMenuId(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindPermissionError, c)
		return
	}
	common.OkWithData(pers, c)
}

// DeletePermission 删除按钮
func (a *PermissionApi) DeletePermission(c *gin.Context) {
	var data system.PermissionId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err := permissionService.DeletePermission(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeletePermissionError, c)
		return
	}
	common.Ok(c)
}
