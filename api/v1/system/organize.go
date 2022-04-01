package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
)

type OrganizeApi struct{}

// Register 注册组织
func (a *OrganizeApi) Register(c *gin.Context) {
	var data system.Organize
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := organizeService.Register(&system.SysOrganize{
		Name: data.Name,
		Pid:  data.Pid,
		Sort: data.Sort,
	}); err != nil {
		common.FailWhitStatus(utils.CreateOrganizationError, c)
		return
	}
	common.Ok(c)
}

// UpdateOrgBaseInfo 更新组织基本信息
func (a *OrganizeApi) UpdateOrgBaseInfo(c *gin.Context) {
	var data system.OrganizeBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := organizeService.Update(&system.SysOrganize{
		BaseID: system.BaseID{
			ID: data.Id,
		},
		Name: data.Name,
		Pid:  data.Pid,
		Sort: data.Sort,
	}); err != nil {
		common.FailWhitStatus(utils.UpdateOrgBaseError, c)
		return
	}
	common.Ok(c)
}

// SearchOrganize 获取组织
func (a *OrganizeApi) SearchOrganize(c *gin.Context) {
	var data system.SearchOrganize
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err, list, total := organizeService.Search(&data)
	if err != nil {
		common.FailWhitStatus(utils.FindOrgError, c)
		return
	}
	common.OkWithData(common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

// DeleteOrganize 删除组织
func (a *OrganizeApi) DeleteOrganize(c *gin.Context) {
	var data system.OrganizeId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err := organizeService.DeleteOrganize(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeleteOrganizationError, c)
		return
	}
	common.Ok(c)
}
