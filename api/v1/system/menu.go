package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
)

type MenuApi struct{}

// Register 注册菜单
func (a *MenuApi) Register(c *gin.Context) {
	var data system.MenuBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := menuService.Register(&system.SysMenu{
		Name:      data.Name,
		Path:      data.Path,
		Hidden:    *data.Hidden,
		Component: data.Component,
		Pid:       data.Pid,
		Sort:      data.Sort,
		Mate: system.Mate{
			Icon: data.Icon,
		},
	}); err != nil {
		common.FailWhitStatus(utils.CreateMenuError, c)
		return
	}
	common.Ok(c)
}

// UpdateMenuBaseInfo 更新菜单基本信息
func (a *MenuApi) UpdateMenuBaseInfo(c *gin.Context) {
	var data system.MenuBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := menuService.Update(system.SysMenu{
		BaseID: system.BaseID{
			ID: data.Id,
		},
		Name:      data.Name,
		Title:     data.Title,
		Path:      data.Path,
		Hidden:    *data.Hidden,
		Component: data.Component,
		Pid:       data.Pid,
		Sort:      data.Sort,
		Mate: system.Mate{
			Icon: data.Icon,
		},
	}); err != nil {
		common.FailWhitStatus(utils.UpdateMenuBaseError, c)
		return
	}
	common.Ok(c)
}

// MenuAll 所有页面
func (a *MenuApi) MenuAll(c *gin.Context) {
	err, menus := menuService.GetAll()
	if err != nil {
		common.FailWhitStatus(utils.FindMenuError, c)
		return
	}
	common.OkWithData(menus, c)
}

// DeleteMenu 删除菜单
func (a *MenuApi) DeleteMenu(c *gin.Context) {
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
	err := menuService.DeleteMenu(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeleteMenuError, c)
		return
	}
	common.Ok(c)
}
