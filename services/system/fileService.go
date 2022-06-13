package system

import (
	"github.com/google/uuid"
	"permissions/global"
	"permissions/model/system"
)

type FileService struct{}

// Register 注册文件
func (s *FileService) Register(dto any) (err error) {
	err = global.Db.Create(&dto).Error
	return
}

// FileById 获取文件信息
func (s *FileService) FileById(id uuid.UUID) (data system.SysFile, err error) {
	err = global.Db.First(&data, id).Error
	return
}
