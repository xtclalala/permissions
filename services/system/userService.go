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

// Register 注册用户
func (s *UserService) Register(dto *system.SysUser) (err error) {
	if errors.Is(s.CheckRepeat(dto.LoginName), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// UpdateUserInfo 更新用户信息
func (s *UserService) UpdateUserInfo(dto system.SysUser) (err error) {
	var old system.SysUser
	err = global.Db.First(&old, dto.ID).Error
	if err != nil {
		return errors.New("主键查找错误")
	}
	if old.LoginName != dto.LoginName {
		if errors.Is(s.CheckRepeat(dto.LoginName), gorm.ErrRecordNotFound) {
			return errors.New("已被注册")
		}
	}
	dto.Password = old.Password
	err = global.Db.Updates(&dto).Error
	return
}

// ResetPassword 重置密码
func (s *UserService) ResetPassword(id uuid.UUID) (err error) {
	var old system.SysUser
	err = global.Db.First(&old, id).Update("password", "123456@y1t").Error
	if err != nil {
		return err
	}
	return
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(id *uuid.UUID, password *string) (err error) {
	var old system.SysUser
	err = global.Db.First(&old, id).Update("password", password).Error
	if err != nil {
		return err
	}
	return
}

// SetUserRoleAndOrg 设置用户权限 角色 组织
func (s *UserService) SetUserRoleAndOrg(userId uuid.UUID, roleIds []int, orgIds []int) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		// 删除原角色
		if err := tx.Where(&system.M2mUserRole{SysUserId: userId}).Delete(&system.M2mUserRole{}).Error; err != nil {
			return err
		}
		if err := tx.Where(&system.M2mUserOrganize{SysUserId: userId}).Delete(&system.M2mUserOrganize{}).Error; err != nil {
			return err
		}

		var m2mUserOrgs []system.M2mUserOrganize
		var m2mUserRoles []system.M2mUserRole
		for _, orgId := range orgIds {
			m2mUserOrgs = append(m2mUserOrgs, system.M2mUserOrganize{
				SysUserId:     userId,
				SysOrganizeId: orgId,
			})
		}
		for _, roleId := range roleIds {
			m2mUserRoles = append(m2mUserRoles, system.M2mUserRole{
				SysUserId: userId,
				SysRoleId: roleId,
			})
		}
		if err := tx.Create(&m2mUserRoles).Error; err != nil {
			return err
		}
		if err := tx.Create(&m2mUserOrgs).Error; err != nil {
			return err
		}
		return nil
	})
}

// Search 搜索用户
func (s *UserService) Search(dto system.SearchUser) (err error, list []system.SysUser, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()

	db := global.Db.Model(&system.SysUser{})

	if len(dto.SysOrganizeIds) != 0 {
		var orgs []system.M2mUserOrganize
		global.Db.Model(&system.M2mUserOrganize{}).Where("sys_organize_id in ?", dto.SysOrganizeIds).Find(&orgs)
		var ids []uuid.UUID
		for _, org := range orgs {
			ids = append(ids, org.SysUserId)
		}
		db = db.Where("id in ?", ids)
	}
	if len(dto.SysRoleIds) != 0 {
		var roles []system.M2mUserRole
		global.Db.Model(&system.M2mUserRole{}).Where("sys_role_id in ?", dto.SysRoleIds).Find(&roles)
		var ids []uuid.UUID
		for _, role := range roles {
			ids = append(ids, role.SysUserId)
		}
		db = db.Where("id in ?", ids)
	}
	if dto.Username != "" {
		db = db.Where("username like ?", "%"+dto.Username+"%")
	}
	if dto.LoginName != "" {
		db = db.Where("login_name like ?", "%"+dto.LoginName+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Limit(limit).Offset(offset)

	oc := clause.OrderByColumn{
		Column: clause.Column{Name: "create_time", Raw: true},
		Desc:   dto.Desc,
	}
	//err = db.Preload(clause.Associations).Select([]string{"username", "id", "LoginName"}).Order(oc).Find(&list).Error
	err = db.Preload(clause.Associations).Omit("password").Order(oc).Find(&list).Error
	return
}

// CheckRepeat 检查 账号 是否存在
func (s *UserService) CheckRepeat(loginName string) (err error) {
	var total int64
	global.Db.Model(&system.SysUser{}).Where(&system.SysUser{LoginName: loginName}).Count(&total)
	if total != 0 {
		err = gorm.ErrRecordNotFound
	} else {
		err = nil
	}
	return
}

// GetAll 查所有用户
func (s *UserService) GetAll() (err error, dos []system.SysUser) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查用户
func (s *UserService) GetById(id uuid.UUID) (err error, do system.SysUser) {
	err = global.Db.First(&do, id).Error
	return
}

// GetCompleteInfoById 根据 id 查用户完整信息
func (s *UserService) GetCompleteInfoById(id uuid.UUID) (err error, do system.SysUser) {
	err = global.Db.Preload(clause.Associations).First(&do, id).Error
	return
}

// GetUserByRoleId 根据 角色id 查用户
func (s *UserService) GetUserByRoleId(roleId int) (err error, users []system.SysUser) {
	rows, err := global.Db.Model(system.M2mUserRole{}).Where(&system.M2mUserRole{SysRoleId: roleId}).Rows()
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
func (s *UserService) GetUserByOrgId(orgId int) (err error, users []system.SysUser) {
	rows, err := global.Db.Model(&system.M2mUserOrganize{}).Where(&system.M2mUserOrganize{SysOrganizeId: orgId}).Rows()
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

// GetUserByLoginName 根据账号查用户
func (s *UserService) GetUserByLoginName(loginName string) (error, system.SysUser) {
	var user system.SysUser
	if err := global.Db.Where(&system.SysUser{LoginName: loginName}).First(&user).Error; err != nil {
		return err, user
	}
	return nil, user
}

func (s *UserService) Delete(userId uuid.UUID) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&system.M2mUserRole{SysUserId: userId}).Delete(&system.M2mUserRole{}).Error; err != nil {
			return err
		}
		if err := tx.Where(&system.M2mUserOrganize{SysUserId: userId}).Delete(&system.M2mUserOrganize{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&system.SysUser{}, userId).Error; err != nil {
			return err
		}
		return nil
	})
}
