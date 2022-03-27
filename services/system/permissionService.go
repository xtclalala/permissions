package system

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	system2 "permissions/model/system"
)

type PermissionService struct{}

var AppPermissionService = new(PermissionService)

// Register 注册页面按钮
func (s *PermissionService) Register(dto *system2.SysPermission) (err error) {
	if errors.Is(s.CheckRepeat(dto.SysMenuId, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// Update 更新页面按钮
func (s *PermissionService) Update(dto system2.SysPermission) (err error) {
	var old system2.SysPermission
	err = global.Db.Where("id = ?", dto.ID).Find(&old).Error
	if err != nil {
		return errors.New("主键查找错误")
	}
	if old.SysMenuId != dto.SysMenuId || old.Name != dto.Name {
		if errors.Is(s.CheckRepeat(dto.SysMenuId, dto.Name), gorm.ErrRecordNotFound) {
			return errors.New("已被注册")
		}
	}

	err = global.Db.Save(dto).Error
	return
}

// Search 搜索菜单
func (s *PermissionService) Search(dto system2.SearchPermission) (err error, list []system2.SysPermission, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system2.SysPermission{})
	var menus []system2.SysPermission

	if dto.SysMenuId != 0 {
		db = db.Where("sys_menu_id = ?", dto.SysMenuId)
	}
	if dto.Name != "" {
		db = db.Where("name like ?", "%"+dto.Name+"%")
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

// CheckRepeat 检查 页面下的按钮 是否存在
func (s *PermissionService) CheckRepeat(menuId int, name string) (err error) {
	var temp system2.SysPermission
	err = global.Db.Where("sys_menu_id = ? and name = ?", menuId, name).First(&temp).Error
	return
}

// GetAll 查所有页面按钮
func (s *PermissionService) GetAll() (err error, dos []system2.SysPermission) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查 按钮
func (s *PermissionService) GetById(id int) (err error, do system2.SysPermission) {
	err = global.Db.Where("id = ?", id).First(&do).Error
	return
}

// GetPerByMenuId 根据 菜单id 查 按钮
func (s *PermissionService) GetPerByMenuId(menuId int) (err error, pers []system2.SysPermission) {
	err = global.Db.Where("sys_menu_id = ?", menuId).Find(&pers).Error
	return
}

// GetPerByRole 根据 角色id 查按钮
func (s *PermissionService) GetPerByRoleId(roleId int) (err error, pers []system2.SysPermission) {
	rows, err := global.Db.Table("m2m_role_permission").Where("sys_role_id = ?", roleId).Rows()
	defer rows.Close()
	if err != nil {
		return err, pers
	}
	for rows.Next() {
		var rolePer system2.M2mRolePermission
		global.Db.ScanRows(rows, &rolePer)
		_, per := s.GetById(rolePer.SysPermissionId)
		pers = append(pers, per)
	}
	return
}

// Ids2Object id转对象
func (s *PermissionService) Ids2Object(ids []int) (pers []system2.SysPermission) {
	for _, id := range ids {
		pers = append(pers, system2.SysPermission{BaseID: system2.BaseID{ID: id}})
	}
	return
}

// DeletePermission 删除按钮
func (s *PermissionService) DeletePermission(id int) (err error) {
	err = global.Db.Where("id = ?", id).Delete(&system2.SysPermission{}).Error
	return
}
