package system

import (
	"errors"
	"gorm.io/gorm"
	"permissions/global"
	"permissions/model/system"
)

type MenuService struct{}

// Register 注册页面
func (s *MenuService) Register(dto system.SysMenu) (err error, do system.SysMenu) {
	var menuTemp system.SysMenu
	if errors.Is(global.Db.Where(" path = ? ", dto.Path).Find(&menuTemp).Error, gorm.ErrRecordNotFound) {
		return errors.New("path已被注册"), menuTemp
	}
	err = global.Db.Create(&dto).Error
	return

}
