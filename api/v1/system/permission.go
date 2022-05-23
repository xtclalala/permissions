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
	err := utils.Validate(&data)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := permissionService.Register(&system.SysPermission{
		Name:      data.Name,
		Code:      data.Code,
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
	err := utils.Validate(&data)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := permissionService.Update(&system.SysPermission{
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
	err := utils.Validate(&data)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, pers := permissionService.GetPerByMenuId(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindPermissionError, c)
		return
	}
	common.OkWithData(pers, c)
}

// SearchPermission 搜索页面权限
func (a *PermissionApi) SearchPermission(c *gin.Context) {
	var data system.SearchPermission
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	err := utils.Validate(&data)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, list, total := permissionService.Search(&data)
	if err != nil {
		common.FailWhitStatus(utils.FindOrgError, c)
		return
	}
	common.OkWithData(common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

// DeletePermission 删除按钮
func (a *PermissionApi) DeletePermission(c *gin.Context) {
	var data system.PermissionId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	err := utils.Validate(&data)
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err = permissionService.DeletePermission(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeletePermissionError, c)
		return
	}
	common.Ok(c)
}
