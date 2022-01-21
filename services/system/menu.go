package system

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	"permissions/model/system"
	"permissions/model/system/search/request"
)

type MenuService struct{}

// Register 注册页面
func (s *MenuService) Register(dto system.SysMenu) (err error) {
	if errors.Is(s.CheckRepeat(dto.Path, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("path已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// Update 更新页面
func (s *MenuService) Update(dto system.SysMenu) (err error) {
	if errors.Is(s.CheckRepeat(dto.Path, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("path已被注册")
	}
	err = global.Db.Model(&dto).Error
	return
}

// GetAll 查所有页面
func (s *MenuService) GetAll() (err error, dos []system.SysMenu) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查页面
func (s *MenuService) GetById(id uint) (err error, do system.SysMenu) {
	err = global.Db.Where("id = ?", id).First(&do).Error
	return
}

// Search 搜索菜单
func (s *MenuService) Search(dto request.SearchMenu) (err error, list []system.SysMenu, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system.SysMenu{})
	var menus []system.SysMenu

	if dto.Path != "" {
		db = db.Where("path like ?", "%"+dto.Path+"%")
	}
	if dto.Name != "" {
		db = db.Where("name like ?", "%"+dto.Name+"%")
	}
	if &(dto.Hidden) != nil {
		db = db.Where("hidden = ?", dto.Hidden)
	}

	err = db.Count(&total).Error
	if err != nil {
		return err, menus, total
	}
	db = db.Limit(limit).Offset(offset)

	oc := clause.OrderByColumn{
		Desc: dto.Desc,
	}
	if dto.Order != "" {
		oc.Column = clause.Column{Name: dto.Order, Raw: true}
	}
	err = db.Order(oc).Find(&list).Error
	return err, list, total
}

// CheckRepeat 检查path 或
func (s *MenuService) CheckRepeat(path, name string) (err error) {
	var total int64
	global.Db.Where("path = ? or name = ?", path, name).Count(&total)
	if total != 0 {
		err = gorm.ErrRecordNotFound
	} else {
		err = nil
	}
	return
}
