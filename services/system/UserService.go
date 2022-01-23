package system

import (
	"errors"
	"github.com/google/uuid"
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

// UpdateUserInfo 更新用户信息
func (s *UserService) UpdateUserInfo(dto system.SysUser) (err error) {
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

// SetRole 设置用户权限 角色
func (s *UserService) SetUserRole(userId uuid.UUID, roleIds []uint) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 删除原角色
		if err := tx.Where(&system.M2mUserRole{SysUserId: userId}).Delete(&system.M2mUserRole{}).Error; err != nil {
			return err
		}
		m2mUserRoles := []system.M2mUserRole{}
		for _, roleId := range roleIds {
			m2mUserRoles = append(m2mUserRoles, system.M2mUserRole{
				SysUserId: userId,
				SysRoleId: roleId,
			})
		}
		if err := tx.Create(&m2mUserRoles).Error; err != nil {
			return err
		}
		return nil
	})
}

// SetOrg 设置用户权限 组织
func (s *UserService) SetUserOrg(userId uuid.UUID, OrgIds []uint) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 删除原组织
		if err := tx.Where(&system.M2mUserOrganize{SysUserId: userId}).Delete(&system.M2mUserOrganize{}).Error; err != nil {
			return err
		}
		m2mUserOrgs := []system.M2mUserOrganize{}
		for _, orgId := range OrgIds {
			m2mUserOrgs = append(m2mUserOrgs, system.M2mUserOrganize{
				SysUserId:     userId,
				SysOrganizeId: orgId,
			})
		}
		if err := tx.Create(&m2mUserOrgs).Error; err != nil {
			return err
		}
		return nil
	})
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
func (s *UserService) GetById(id uuid.UUID) (err error, do system.SysUser) {
	err = global.Db.Where("id = ?", id).First(&do).Error
	return
}

// GetUserByRoleId 根据 角色id 查用户
func (s *UserService) GetUserByRoleId(roleId uint) (err error, users []system.SysUser) {
	rows, err := global.Db.Where(&system.M2mUserRole{SysRoleId: roleId}).Rows()
	defer rows.Close()
	if err != nil {
		return err, users
	}
	for rows.Next() {
		var userRole system.M2mUserRole
		global.Db.ScanRows(rows, &userRole)
		_, user := s.GetById(userRole.SysUserId)
		users = append(users, user)
	}
	return
}

// GetUserByOrgId 根据 组织id 查用户
func (s *UserService) GetUserByOrgId(orgId uint) (err error, users []system.SysUser) {
	rows, err := global.Db.Where(&system.M2mUserOrganize{SysOrganizeId: orgId}).Rows()
	defer rows.Close()
	if err != nil {
		return err, users
	}
	for rows.Next() {
		var userRole system.M2mUserRole
		global.Db.ScanRows(rows, &userRole)
		_, user := s.GetById(userRole.SysUserId)
		users = append(users, user)
	}
	return
}
