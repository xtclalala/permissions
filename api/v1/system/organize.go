package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	system2 "permissions/model/system"
	utils2 "permissions/utils"
)

type OrganizeApi struct{}

// Register 注册组织
func (a *OrganizeApi) Register(c *gin.Context) {
	var data system2.Organize
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := organizeService.Register(&system2.SysOrganize{
		Name: data.Name,
		Pid:  data.Pid,
		Sort: data.Sort,
	}); err != nil {
		common.FailWhitStatus(utils2.CreateOrganizationError, c)
		return
	}
	common.Ok(c)
}

// UpdateOrgBaseInfo 更新组织基本信息
func (a *OrganizeApi) UpdateOrgBaseInfo(c *gin.Context) {
	var data system2.OrganizeBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := organizeService.Update(&system2.SysOrganize{
		BaseID: system2.BaseID{
			ID: data.Id,
		},
		Name: data.Name,
		Pid:  data.Pid,
		Sort: data.Sort,
	}); err != nil {
		common.FailWhitStatus(utils2.UpdateOrgBaseError, c)
		return
	}
	common.Ok(c)
}

// SearchOrganize 获取组织
func (a *OrganizeApi) SearchOrganize(c *gin.Context) {
	var data system2.SearchOrganize
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err, list, total := organizeService.Search(&data)
	if err != nil {
		common.FailWhitStatus(utils2.FindOrgError, c)
		return
	}
	common.OkWithData(common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

// DeleteOrganize 删除组织
func (a *OrganizeApi) DeleteOrganize(c *gin.Context) {
	var data system2.OrganizeId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err := organizeService.DeleteOrganize(data.Id)
	if err != nil {
		common.FailWhitStatus(utils2.DeleteOrganizationError, c)
		return
	}
	common.Ok(c)
}
