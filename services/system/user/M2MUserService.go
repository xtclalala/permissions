package user

import (
	"permissions/global"
	"permissions/model/system"
	//"permissions/services/system/organize"
)

type M2MUserService struct{}

var AppM2MUserService = new(M2MUserService)

//Update 修改用户权限 角色，组织
func (s *M2MUserService) Update(dto system.SysUser) (err error) {
	err = s.CheckRoleAndOrg(dto.SysRoles, dto.SysOrganizes)
	if err != nil {
		return
	}
	err = global.Db.Model(&dto).Association("SysRoles").Replace(dto.SysRoles)
	if err != nil {
		return
	}
	err = global.Db.Model(&dto).Association("SysOrganizes").Replace(dto.SysOrganizes)
	return
}

// CheckRoleAndOrg 检查 角色是否属于组织
func (s *M2MUserService) CheckRoleAndOrg(roles []system.SysRole, organizes []system.SysOrganize) (err error) {
	//for  _, organize := range organizes {
	// todo 查组织和角色表 该角色是否是该组织下的 不清楚要不要
	//}
	//err = global.Db.Where("login_name = ?", loginName).First(&temp).Error
	return
}

// GetOrgByUser 根据 用户id 查组织
func (s *M2MUserService) GetOrgByUser(user system.SysUser) (err error, orgs []system.SysOrganize) {
	err = global.Db.Model(&user).Association("SysOrganizes").Find(&orgs)
	return
}

// GetRoleByUser 根据 用户id 查角色
func (s *M2MUserService) GetRoleByUser(user system.SysUser) (err error, roles []system.SysRole) {
	err = global.Db.Model(&user).Association("SysRoles").Find(&roles)
	return
}

// GetUserByRole 根据 角色id 查用户
func (s *M2MUserService) GetUserByRole(role system.SysRole) (err error, users []system.SysUser) {
	err = global.Db.Model(&role).Association("SysUsers").Find(&users)
	return
}

// GetUserByOrg 根据 组织id 查用户
func (s *M2MUserService) GetUserByOrg(org system.SysOrganize) (err error, users []system.SysUser) {
	err = global.Db.Model(&org).Association("SysUsers").Find(&users)
	return
}
