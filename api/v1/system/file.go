package system

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"path"
	"permissions/global"
	"permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
)

type FileApi struct{}

// Upload 上传文件
func (a *FileApi) Upload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]
	var dataList []*system.SysFile
	var dataMap map[string]string
	uploadPath := global.System.File.Path
	for _, file := range files {
		id := uuid.New()
		idStr := id.String()
		b, _ := file.Open()
		defer b.Close()
		fileType, err := utils.GetFileType(b)
		if err != nil {
			common.FailWhitStatus(utils.FileReadType, c)
			return
		}
		f := &system.SysFile{
			BaseUUID: system.BaseUUID{
				ID: id,
			},
			Name: file.Filename,
			Type: fileType,
			Path: uploadPath,
		}
		dataList = append(dataList, f)
		dataMap[idStr] = file.Filename
		file.Filename = idStr
		if err := c.SaveUploadedFile(file, path.Join(uploadPath, file.Filename)); err != nil {
			common.FailWhitStatus(utils.CreateMenuError, c)
			return
		}
	}
	if err := fileService.Register(dataList); err != nil {
		common.FailWhitStatus(utils.CreateMenuError, c)
		return
	}
	common.OkWithData(dataMap, c)
}

// Download 获取文件信息
func (a *FileApi) Download(c *gin.Context) {
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
	c.File(file.Path + utils.Slash + file.Name)
	common.Ok(c)
}
