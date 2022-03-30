package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	system2 "permissions/model/system"
	utils2 "permissions/utils"
)

type MenuApi struct{}

// Register 注册菜单
func (a *MenuApi) Register(c *gin.Context) {
	var data system2.MenuBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := menuService.Register(&system2.SysMenu{
		Name:      data.Name,
		Path:      data.Path,
		Hidden:    *data.Hidden,
		Component: data.Component,
		Pid:       data.Pid,
		Sort:      data.Sort,
		Mate: system2.Mate{
			Icon: data.Icon,
		},
	}); err != nil {
		common.FailWhitStatus(utils2.CreateMenuError, c)
		return
	}
	common.Ok(c)
}

// UpdateMenuBaseInfo 更新菜单基本信息
func (a *MenuApi) UpdateMenuBaseInfo(c *gin.Context) {
	var data system2.MenuBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := menuService.Update(system2.SysMenu{
		BaseID: system2.BaseID{
			ID: data.Id,
		},
		Name:      data.Name,
		Path:      data.Path,
		Hidden:    *data.Hidden,
		Component: data.Component,
		Pid:       data.Pid,
		Sort:      data.Sort,
		Mate: system2.Mate{
			Icon: data.Icon,
		},
	}); err != nil {
		common.FailWhitStatus(utils2.UpdateMenuBaseError, c)
		return
	}
	common.Ok(c)
}

// MenuAll 所有页面
func (a *MenuApi) MenuAll(c *gin.Context) {
	err, menus := menuService.GetAll()
	if err != nil {
		common.FailWhitStatus(utils2.FindMenuError, c)
		return
	}
	common.OkWithData(menus, c)
}

// DeleteMenu 删除菜单
func (a *MenuApi) DeleteMenu(c *gin.Context) {
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
	err := menuService.DeleteMenu(data.Id)
	if err != nil {
		common.FailWhitStatus(utils2.DeleteMenuError, c)
		return
	}
	common.Ok(c)
}
