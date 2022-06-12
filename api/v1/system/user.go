package system

import (
	"github.com/gin-gonic/gin"
	"permissions/global"
	"permissions/model/common"
	"permissions/model/system"
	services "permissions/services/system"
	"permissions/utils"
)

type UserApi struct{}

// CreateUser 创建用户
func (a *UserApi) CreateUser(c *gin.Context) {
	var data system.User
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	var roles []system.SysRole
	var orgs []system.SysOrganize
	for _, roleId := range data.SysRoleIds {
		roles = append(roles, system.SysRole{BaseID: system.BaseID{ID: roleId}})
	}
	for _, orgId := range data.SysOrganizeIds {
		orgs = append(orgs, system.SysOrganize{BaseID: system.BaseID{ID: orgId}})
	}
	if err := userService.Register(&system.SysUser{
		Username:     data.Username,
		LoginName:    data.LoginName,
		Password:     data.LoginName + "@y1t",
		SysRoles:     roles,
		SysOrganizes: orgs,
	}); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}

// UpdateUserBaseInfo 更新基本用户
func (a *UserApi) UpdateUserBaseInfo(c *gin.Context) {
	var data system.User
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.SetUserRoleAndOrg(data.Id, data.SysRoleIds, data.SysOrganizeIds); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.UpdateUserInfo(system.SysUser{
		BaseUUID: system.BaseUUID{
			ID: data.Id,
		},
		Username:  data.Username,
		LoginName: data.LoginName,
	}); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}

// UpdateUserPerInfo 更新用户权限
func (a *UserApi) UpdateUserPerInfo(c *gin.Context) {
	var data system.UserPerInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.SetUserRoleAndOrg(data.Id, data.SysRoleIds, data.SysOrganizeIds); err != nil {
		common.Fail(c)
		return
	}
	common.Ok(c)
}

// Login 登录
func (a *UserApi) Login(c *gin.Context) {
	var data system.UserLogin
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, user := userService.GetUserByLoginName(data.LoginName)
	if err != nil {
		common.FailWhitStatus(utils.UsernameAndPasswdError, c)
		return
	}
	if user.Password != data.Password {
		common.FailWhitStatus(utils.UsernameAndPasswdError, c)
		return
	}
	j := services.NewJWT()
	claim := j.CreateClaim(&user)
	token, err := j.CreateJwt(&claim)
	if err != nil {
		common.FailWhitStatus(utils.TokenCreateFailed, c)
		return
	}
	common.OkWithData(token, c)
}

// GetUserRouterAndRoles 当前登录用户完整信息
func (a *UserApi) GetUserRouterAndRoles(c *gin.Context) {
	token := c.GetHeader(global.System.App.Auth)
	j := services.NewJWT()
	claim, err := j.ParseJwt(token)
	if err != 0 {
		common.FailWhitStatus(err, c)
		return
	}
	//claim, _ := c.Get("claims")
	//claims := claim.(common.Y1tClaim)
	ok, orgs := organizeService.GetOrgByUserId(claim.Id)
	if ok != nil {
		common.FailWhitStatus(utils.FindOrgError, c)
		return
	}
	roles := make([]system.SysRole, 0, 10)
	for _, org := range orgs {
		ok, sysRoles := roleService.GetRoleByOrgId(org.ID)
		if ok != nil {
			common.FailWhitStatus(utils.FindRoleError, c)
			return
		}

		for i, role := range sysRoles {
			pers := make([]system.SysPermission, 0, 50)
			ok, SysPermissions := permissionService.GetPerByRoleId(role.ID)
			if ok != nil {
				common.FailWhitStatus(utils.FindPermissionError, c)
				return
			}
			pers = append(pers, SysPermissions...)
			sysRoles[i].SysPermissions = pers
		}

		for i, role := range sysRoles {
			menus := make([]system.SysMenu, 0, 30)
			ok, sysMenus := menuService.GetMenuByRoleId(role.ID)
			if ok != nil {
				common.FailWhitStatus(utils.FindMenuError, c)
				return
			}
			menus = append(menus, sysMenus...)
			sysRoles[i].SysMenus = menus
		}

		roles = append(roles, sysRoles...)
	}

	data := map[string]any{
		"username": claim.Username,
		"roles":    roles,
		"orgs":     orgs,
	}
	common.OkWithData(data, c)
}

// SearchUsers 搜索
func (a *UserApi) SearchUsers(c *gin.Context) {
	var data system.SearchUser
	if err := c.ShouldBindQuery(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, list, total := userService.Search(data)
	if err != nil {
		common.FailWhitStatus(utils.FindUserError, c)
		return
	}
	common.OkWithData(common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

// CompleteInfo 用户详细信息
func (a *UserApi) CompleteInfo(c *gin.Context) {
	var data system.UserId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	err, user := userService.GetCompleteInfoById(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindUserError, c)
		return
	}
	common.OkWithData(user, c)
}

// DeleteUser 删除用户
func (a *UserApi) DeleteUser(c *gin.Context) {
	var data system.UserId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}

	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.Delete(data.Id); err != nil {
		common.FailWhitStatus(utils.DeleteUserError, c)
		return
	}
	common.Ok(c)
}

// ResetPassword 重置密码
func (a *UserApi) ResetPassword(c *gin.Context) {
	var data system.UserId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
		return
	}
	if err := utils.Validate(&data); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	if err := userService.ResetPassword(data.Id); err != nil {
		common.FailWhitStatus(utils.DeleteUserError, c)
		return
	}
	common.Ok(c)
}
