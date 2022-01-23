package menu

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	"permissions/model/system"
)

type MenuService struct{}

var AppMenuService = new(MenuService)

// Register 注册页面
func (s *MenuService) Register(dto system.SysMenu) (err error) {
	if errors.Is(s.CheckRepeat(dto.Path, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// Update 更新页面
func (s *MenuService) Update(dto system.SysMenu) (err error) {
	var old system.SysMenu
	err = global.Db.Where("id = ?", dto.ID).Find(&old).Error
	if err != nil {
		return errors.New("主键查找错误")
	}
	if old.Name != dto.Name || old.Path != dto.Path {
		if errors.Is(s.CheckRepeat(dto.Path, dto.Name), gorm.ErrRecordNotFound) {
			return errors.New("已被注册")
		}
	}
	err = global.Db.Save(dto).Error
	return
}

// Search 搜索菜单
func (s *MenuService) Search(dto SearchMenu) (err error, list []system.SysMenu, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system.SysMenu{})
	var menus []system.SysMenu

	if dto.Pid != 0 {
		db = db.Where("pid = ?", dto.Pid)
	}
	if dto.Path != "" {
		db = db.Where("path like ?", "%"+dto.Path+"%")
	}
	if dto.Name != "" {
		db = db.Where("name like ?", "%"+dto.Name+"%")
	}
	if dto.Component != "" {
		db = db.Where("component like ?", "%"+dto.Component+"%")
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
		Column: clause.Column{Name: "sort", Raw: true},
		Desc:   dto.Desc,
	}

	err = db.Order(oc).Find(&list).Error
	return err, list, total
}

// CheckRepeat 检查path 或 name 是否存在
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

// GetMenuByRole 根据 角色id 查菜单
func (s *MenuService) GetMenuByRoleId(roleId uint) (err error, menus []system.SysMenu) {
	rows, err := global.Db.Where(&system.M2mRoleMenu{SysRoleId: roleId}).Rows()
	defer rows.Close()
	if err != nil {
		return err, menus
	}
	for rows.Next() {
		var roleMenu system.M2mRoleMenu
		global.Db.ScanRows(rows, &roleMenu)
		_, menu := s.GetById(roleMenu.SysMenuId)
		menus = append(menus, menu)
	}
	return
}
