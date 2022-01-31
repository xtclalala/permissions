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
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := organizeService.Register(&system.SysOrganize{
		Name: data.Name,
		Pid:  data.Pid,
		Sort: data.Sort,
	}); err != nil {
		common.FailWhitStatus(utils.CreateOrganizationError, c)
	}
	common.Ok(c)
}

// UpdateOrgBaseInfo 更新组织基本信息
func (a *OrganizeApi) UpdateOrgBaseInfo(c *gin.Context) {
	var data system.OrganizeBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := organizeService.Update(&system.SysOrganize{
		Name: data.Name,
		Pid:  data.Pid,
		Sort: data.Sort,
	}); err != nil {
		common.FailWhitStatus(utils.UpdateOrgBaseError, c)
	}
	common.Ok(c)
}

// SearchOrganize 获取组织
func (a *OrganizeApi) SearchOrganize(c *gin.Context) {
	var data system.SearchOrganize
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, list, total := organizeService.Search(&data)
	if err != nil {
		common.FailWhitStatus(utils.FindOrgError, c)
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
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err := organizeService.DeleteOrganize(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.DeleteOrganizationError, c)
	}
	common.Ok(c)
}
