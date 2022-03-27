package system

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"permissions/global"
	system2 "permissions/model/system"
)

type RoleService struct{}

var AppRoleService = new(RoleService)

// Register 注册角色
func (s *RoleService) Register(dto system2.SysRole) (err error) {
	if errors.Is(s.CheckRepeat(dto.Pid, dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// UpdateRoleInfo 更新角色
func (s *RoleService) UpdateRoleInfo(dto system2.SysRole) (err error) {
	var old system2.SysRole
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

// SetRoleMenu 修改角色权限 菜单
func (s *RoleService) SetRoleMenu(roleId int, menuIds []int) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&system2.M2mRoleMenu{SysRoleId: roleId}).Delete(&system2.M2mRoleMenu{}).Error; err != nil {
			return err
		}
		var roleMenus []system2.M2mRoleMenu
		for _, menuId := range menuIds {
			roleMenus = append(roleMenus, system2.M2mRoleMenu{
				SysRoleId: roleId,
				SysMenuId: menuId,
			})
		}
		if err := tx.Create(&roleMenus).Error; err != nil {
			return err
		}
		return nil
	})

}

// SetRolePer 修改角色权限 按钮
func (s *RoleService) SetRolePer(roleId int, perIds []int) (err error) {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&system2.M2mRolePermission{SysRoleId: roleId}).Delete(&system2.M2mRolePermission{}).Error; err != nil {
			return err
		}
		var rolePers []system2.M2mRolePermission
		for _, perId := range perIds {
			rolePers = append(rolePers, system2.M2mRolePermission{
				SysRoleId:       roleId,
				SysPermissionId: perId,
			})
		}
		if err := tx.Create(&rolePers).Error; err != nil {
			return err
		}
		return nil
	})
}

// Search 搜索角色
func (s *RoleService) Search(dto system2.SearchRole) (err error, list []system2.SysRole, total int64) {
	limit := dto.PageSize
	offset := dto.GetOffset()
	db := global.Db.Model(&system2.SysRole{})
	var menus []system2.SysRole

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
func (s *RoleService) CheckRepeat(pid int, name string) (err error) {
	var temp system2.SysRole
	err = global.Db.Where("pid = ? and name = ?", pid, name).First(&temp).Error
	return
}

// GetAll 查所有角色
func (s *RoleService) GetAll() (err error, dos []system2.SysRole) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查角色
func (s *RoleService) GetById(id int) (err error, do system2.SysRole) {
	err = global.Db.Where("id = ?", id).First(&do).Error
	return
}

// GetById 根据 id 查角色
func (s *RoleService) GetCompleteInfoById(id int) (err error, do system2.SysRole) {
	err = global.Db.Preload(clause.Associations).Where("id = ?", id).First(&do).Error
	return
}

// GetRoleByUserId 根据 用户id 查角色
func (s *RoleService) GetRoleByUserId(userId uuid.UUID) (err error, roles []system2.SysRole) {
	rows, err := global.Db.Where(&system2.M2mUserRole{}, userId).Rows()
	defer rows.Close()
	if err != nil {
		return err, roles
	}
	for rows.Next() {
		var userRole system2.M2mUserRole
		global.Db.ScanRows(rows, &userRole)
		_, role := s.GetById(userRole.SysRoleId)
		roles = append(roles, role)
	}
	return
}

// GetRoleByMenu 根据 菜单id 查角色
func (s *RoleService) GetRoleByMenuId(menuId int) (err error, roles []system2.SysRole) {
	rows, err := global.Db.Where(&system2.M2mRoleMenu{SysMenuId: menuId}).Rows()
	defer rows.Close()
	if err != nil {
		return err, roles
	}
	for rows.Next() {
		var roleMenu system2.M2mRoleMenu
		global.Db.ScanRows(rows, &roleMenu)
		_, role := s.GetById(roleMenu.SysRoleId)
		roles = append(roles, role)
	}
	return
}

// GetRoleByPer 根据 按钮id 查角色
func (s *RoleService) GetRoleByPerId(perId int) (err error, roles []system2.SysRole) {
	rows, err := global.Db.Where(&system2.M2mRolePermission{SysPermissionId: perId}).Rows()
	defer rows.Close()
	if err != nil {
		return err, roles
	}
	for rows.Next() {
		var rolePer system2.M2mRolePermission
		global.Db.ScanRows(rows, rolePer)
		_, role := s.GetById(rolePer.SysRoleId)
		roles = append(roles, role)
	}
	return
}

// GetRoleByOrgId 根据 组织id 查角色
func (s *RoleService) GetRoleByOrgId(orgId int) (err error, roles []system2.SysRole) {
	err = global.Db.Where(&system2.SysRole{SysOrganizeId: orgId}).Find(&roles).Error
	return
}

func (s *RoleService) DeleteRole(roleId int) (err error) {
	err = global.Db.Where("id = ?", roleId).Delete(&system2.SysRole{}).Error
	return
}
