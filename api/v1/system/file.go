package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
)

type FileApi struct{}

// Upload 上传文件
func (a *MenuApi) Upload(c *gin.Context) {
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

// Download 获取文件信息
func (a *MenuApi) Download(c *gin.Context) {
	var data system.File
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	file, err := fileService.FileById(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.UpdateMenuBaseError, c)
		return
	}
	// 读文件
	// do something
	common.OkWithData(file, c)
}
