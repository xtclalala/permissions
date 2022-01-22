package role

import (
	"permissions/global"
	"permissions/model/system"
)

type M2MRoleService struct{}

var AppM2MRoleService = new(M2MRoleService)

//Update 修改角色权限 组织，菜单，按钮
func (s *M2MRoleService) Update(dto system.SysRole) (err error) {
	err = s.CheckMenuAndPer(dto.SysMenus, dto.SysPermissions)
	if err != nil {
		return
	}
	err = global.Db.Model(&dto).Association("SysMenus").Replace(dto.SysMenus)
	if err != nil {
		return
	}
	err = global.Db.Model(&dto).Association("SysPermissions").Replace(dto.SysPermissions)
	return
}

// CheckMenuAndPer 检查 按钮是否属于页面
func (s *M2MRoleService) CheckMenuAndPer(menus []system.SysMenu, permissions []system.SysPermission) (err error) {
	//for  _, organize := range organizes {
	// todo 查菜单表 按钮是不是目标菜单下的
	//}
	//err = global.Db.Where("login_name = ?", loginName).First(&temp).Error
	return
}

// GetMenuByRole 根据 角色id 查菜单
func (s *M2MRoleService) GetMenuByRole(role system.SysRole) (err error, menus []system.SysMenu) {
	err = global.Db.Model(&role).Association("SysMenus").Find(&menus)
	return
}

// GetPerByRole 根据 角色id 查按钮
func (s *M2MRoleService) GetPerByRole(role system.SysRole) (err error, pers []system.SysPermission) {
	err = global.Db.Model(&role).Association("SysPermissions").Find(&pers)
	return
}

// GetOrgByRole 根据 角色id 查组织
func (s *M2MRoleService) GetOrgByRole(role system.SysRole) (err error, org system.SysOrganize) {
	err = global.Db.Model(&role).Association("SysOrganizes").Find(&org)
	return
}

// GetRoleByMenu 根据 菜单id 查角色
func (s *M2MRoleService) GetRoleByMenu(menu system.SysMenu) (err error, roles []system.SysRole) {
	err = global.Db.Model(&menu).Association("SysRoles").Find(&roles)
	return
}

// GetRoleByPer 根据 按钮id 查角色
func (s *M2MRoleService) GetRoleByPer(per system.SysPermission) (err error, roles []system.SysRole) {
	err = global.Db.Model(&per).Association("SysRoles").Find(&roles)
	return
}

// GetRoleByOrgId 根据 组织id 查角色
func (s *M2MRoleService) GetRoleByOrgId(orgId uint) (err error, roles []system.SysRole) {
	err = global.Db.Where("sys_organize_id = ?", orgId).Find(&roles).Error
	return
}
