package role

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	"permissions/model/system"
)

type RoleService struct{}

var AppRoleService = new(RoleService)

// Register 注册角色
func (s *RoleService) Register(dto system.SysRole) (err error) {
	if errors.Is(s.CheckRepeat(dto.Pid, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// Update 更新角色
func (s *RoleService) Update(dto system.SysRole) (err error) {
	var old system.SysRole
	err = global.Db.Where("id = ?", dto.ID).Find(&old).Error
	if err != nil {
		return errors.New("主键查找错误")
	}
	if old.Pid != dto.Pid || old.Name != dto.Name {
		if errors.Is(s.CheckRepeat(dto.Pid, dto.Name), gorm.ErrRecordNotFound) {
			return errors.New("已被注册")
		}
	}
	err = global.Db.Save(dto).Error
	return
}

// Search 搜索角色
func (s *RoleService) Search(dto SearchRole) (err error, list []system.SysRole, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system.SysRole{})
	var menus []system.SysRole

	if dto.Pid != 0 {
		db = db.Where("pid = ?", dto.Pid)
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

// CheckRepeat 检查 pid 和 名称 是否存在
func (s *RoleService) CheckRepeat(pid uint, name string) (err error) {
	var temp system.SysRole
	err = global.Db.Where("pid = ? and name = ?", pid, name).First(&temp).Error
	return
}

// GetAll 查所有角色
func (s *RoleService) GetAll() (err error, dos []system.SysRole) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查角色
func (s *RoleService) GetById(id uint) (err error, do system.SysRole) {
	err = global.Db.Where("id = ?", id).First(&do).Error
	return
}
