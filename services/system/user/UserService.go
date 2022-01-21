package user

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	"permissions/model/system"
)

type UserService struct{}

var AppUserService = new(UserService)

// Register 注册用户
func (s *UserService) Register(dto system.SysUser) (err error) {
	if errors.Is(s.CheckRepeat(dto.LoginName), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// Update 更新用户
func (s *UserService) Update(dto system.SysUser) (err error) {
	var old system.SysUser
	err = global.Db.Where("id = ?", dto.ID).Find(&old).Error
	if err != nil {
		return errors.New("主键查找错误")
	}
	if old.LoginName != dto.LoginName {
		if errors.Is(s.CheckRepeat(dto.LoginName), gorm.ErrRecordNotFound) {
			return errors.New("已被注册")
		}
	}
	err = global.Db.Save(dto).Error
	return
}

// Search 搜索用户
func (s *UserService) Search(dto SearchUser) (err error, list []system.SysUser, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system.SysUser{})
	var menus []system.SysUser

	if dto.Username != "" {
		db = db.Where("user_name = ?", "%"+dto.Username+"%")
	}
	if dto.LoginName != "" {
		db = db.Where("name like ?", "%"+dto.LoginName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return err, menus, total
	}
	db = db.Limit(limit).Offset(offset)

	oc := clause.OrderByColumn{
		Column: clause.Column{Name: "create_time", Raw: true},
		Desc:   dto.Desc,
	}

	err = db.Order(oc).Find(&list).Error
	return err, list, total
}

// CheckRepeat 检查 账号 是否存在
func (s *UserService) CheckRepeat(loginName string) (err error) {
	var temp system.SysUser
	err = global.Db.Where("login_name = ?", loginName).First(&temp).Error
	return
}

// GetAll 查所有用户
func (s *UserService) GetAll() (err error, dos []system.SysUser) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查用户
func (s *UserService) GetById(id uint) (err error, do system.SysUser) {
	err = global.Db.Where("id = ?", id).First(&do).Error
	return
}
