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
	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.Register(&system.SysMenu{
		Title:     data.Title,
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

// UpdateMenuInfo 更新菜单
func (a *MenuApi) UpdateMenuInfo(c *gin.Context) {
	var data system.MenuBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
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

// SearchMenu 搜索菜单
func (a *MenuApi) SearchMenu(c *gin.Context) {
	var data system.SearchMenu
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, list, total := menuService.Search(data)
	if err != nil {
		common.FailWhitStatus(utils.FindRoleError, c)
		return
	}
	common.OkWithData(&common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

func (a *MenuApi) AllMenu(c *gin.Context) {
	err, list := menuService.GetAll()
	if err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.OkWithData(list, c)
}

// DeleteMenu 删除菜单
func (a *MenuApi) DeleteMenu(c *gin.Context) {
	var data system.MenuId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.DeleteMenu(data.Id); err != nil {
		common.FailWhitStatus(utils.DeleteMenuError, c)
		return
	}
	common.Ok(c)
}
