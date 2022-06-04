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

// Register 注册角色
func (s *RoleService) Register(dto system2.SysRole) (err error) {
	if errors.Is(s.CheckRepeat(dto.Name), gorm.ErrRecordNotFound) {
		return errors.New("已被注册")
	}
	err = global.Db.Create(&dto).Error
	return
}

// UpdateRoleInfo 更新角色
func (s *RoleService) UpdateRoleInfo(dto system2.SysRole) (err error) {
	var old system2.SysRole
	err = global.Db.First(&old, dto.ID).Error
	if err != nil {
		return errors.New("主键查找错误")
	}
	if old.Name != dto.Name {
		if errors.Is(s.CheckRepeat(dto.Name), gorm.ErrRecordNotFound) {
			return errors.New("已被注册")
		}
	}
	err = global.Db.Save(dto).Error
	return
}

// SetRoleMenu 修改角色权限 菜单
func (s *RoleService) SetRoleMenu(roleId int, menuIds []int) error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&system2.M2mRoleMenu{}, "sys_role_id = ?", roleId).Error; err != nil {
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
		if err := tx.Delete(&system2.M2mRolePermission{}, "sys_role_id = ?", roleId).Error; err != nil {
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
	db := global.Db.Model(&system2.SysRole{}).Preload("SysOrganize")

	if dto.Name != "" {
		db = db.Where("name like ?", "%"+dto.Name+"%")
	}
	if dto.OrganizeId != 0 {
		db = db.Where("sys_organize_id = ?", dto.OrganizeId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	db = db.Limit(limit).Offset(offset)

	oc := clause.OrderByColumn{
		Column: clause.Column{Name: "sort", Raw: true},
		Desc:   dto.Desc,
	}

	err = db.Order(oc).Find(&list).Error
	return
}

// CheckRepeat 检查 pid 和 名称 是否存在
func (s *RoleService) CheckRepeat(name string) (err error) {
	var total int64
	global.Db.Model(&system2.SysRole{}).Where(&system2.SysRole{Name: name}).Count(&total)
	if total != 0 {
		err = gorm.ErrRecordNotFound
	} else {
		err = nil
	}
	return
}

// GetAll 查所有角色
func (s *RoleService) GetAll() (err error, dos []system2.SysRole) {
	err = global.Db.Find(&dos).Error
	return
}

// GetById 根据 id 查角色
func (s *RoleService) GetById(id int) (err error, do system2.SysRole) {
	err = global.Db.First(&do, id).Error
	return
}

// GetCompleteInfoById 根据 id 查角色
func (s *RoleService) GetCompleteInfoById(id int) (err error, do system2.SysRole) {
	err = global.Db.Preload("SysMenus").Preload("SysPermissions").Preload("SysOrganize").First(&do, id).Error
	return
}

// GetRoleByUserId 根据 用户id 查角色
func (s *RoleService) GetRoleByUserId(userId uuid.UUID) (err error, roles []system2.SysRole) {
	rows, err := global.Db.Model(&system2.M2mUserRole{}).Where(&system2.M2mUserRole{SysUserId: userId}).Rows()
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

// GetRoleByMenuId 根据 菜单id 查角色
func (s *RoleService) GetRoleByMenuId(menuId int) (err error, roles []system2.SysRole) {
	rows, err := global.Db.Model(&system2.M2mRoleMenu{}).Where(&system2.M2mRoleMenu{SysMenuId: menuId}).Rows()
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

// GetRoleByPerId 根据 按钮id 查角色
func (s *RoleService) GetRoleByPerId(perId int) (err error, roles []system2.SysRole) {
	rows, err := global.Db.Model(&system2.M2mRolePermission{}).Where(&system2.M2mRolePermission{SysPermissionId: perId}).Rows()
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
	err = global.Db.Preload("SysMenus").Preload("SysPermissions").Where("sys_organize_id = ?", orgId).Find(&roles).Error
	return
}

func (s *RoleService) DeleteRole(roleId int) (err error) {
	err = global.Db.Delete(&system2.SysRole{}, roleId).Error
	return
}
