package system

import (
	"github.com/gin-gonic/gin"
	"permissions/global"
	"permissions/model/common"
	system3 "permissions/model/system"
	system2 "permissions/services/system"
	utils2 "permissions/utils"
)

type UserApi struct{}

// CreateUser 创建用户
func (a *UserApi) CreateUser(c *gin.Context) {
	var data system3.User
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	var roles []system3.SysRole
	var orgs []system3.SysOrganize
	for _, roleId := range data.SysRoleIds {
		roles = append(roles, system3.SysRole{BaseID: system3.BaseID{ID: roleId}})
	}
	for _, orgId := range data.SysOrganizeIds {
		orgs = append(orgs, system3.SysOrganize{BaseID: system3.BaseID{ID: orgId}})
	}
	if err := userService.Register(&system3.SysUser{
		Username:     data.Username,
		LoginName:    data.LoginName,
		Password:     data.Password,
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
	var data system3.UserBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := userService.UpdateUserInfo(system3.SysUser{
		BaseUUID:  system3.BaseUUID{ID: data.Id},
		Username:  data.Username,
		Password:  data.Password,
		LoginName: data.LoginName,
	}); err != nil {
		common.FailWithMessage(err.Error(), c)
		return
	}
	common.Ok(c)
}

// UpdateUserPerInfo 更新用户权限
func (a *UserApi) UpdateUserPerInfo(c *gin.Context) {
	var data system3.UserPerInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
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
	var data system3.UserLogin
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err, user := userService.GetUserByLoginName(data.LoginName)
	if err != nil {
		common.FailWhitStatus(utils2.UsernameAndPasswdError, c)
		return
	}
	if user.Password != data.Password {
		common.FailWhitStatus(utils2.UsernameAndPasswdError, c)
		return
	}
	j := system2.NewJWT()
	claim := j.CreateClaim(&user)
	token, err := j.CreateJwt(&claim)
	if err != nil {
		common.FailWhitStatus(utils2.TokenCreateFailed, c)
		return
	}
	common.OkWithData(token, c)
}

// GetUserRouterAndRoles 当前登录用户完整信息
func (a *UserApi) GetUserRouterAndRoles(c *gin.Context) {
	token := c.GetHeader(global.System.App.Auth)
	j := system2.NewJWT()
	claim, err := j.ParseJwt(token)
	if err != 0 {
		common.FailWhitStatus(err, c)
		return
	}
	ok, orgs := organizeService.GetOrgByUserId(claim.Id)
	if ok != nil {
		common.FailWhitStatus(utils2.FindOrgError, c)
		return
	}
	roles := make([]system3.SysRole, 0, 10)
	for _, org := range orgs {
		ok, sysRoles := roleService.GetRoleByOrgId(org.ID)
		if ok != nil {
			common.FailWhitStatus(utils2.FindRoleError, c)
			return
		}

		for i, role := range sysRoles {
			pers := make([]system3.SysPermission, 0, 50)
			ok, SysPermissions := permissionService.GetPerByRoleId(role.ID)
			if ok != nil {
				common.FailWhitStatus(utils2.FindPermissionError, c)
				return
			}
			pers = append(pers, SysPermissions...)
			sysRoles[i].SysPermissions = pers
		}

		for i, role := range sysRoles {
			menus := make([]system3.SysMenu, 0, 30)
			ok, sysMenus := menuService.GetMenuByRoleId(role.ID)
			if ok != nil {
				common.FailWhitStatus(utils2.FindMenuError, c)
				return
			}
			menus = append(menus, sysMenus...)
			sysRoles[i].SysMenus = menus
		}

		roles = append(roles, sysRoles...)
	}

	data := map[string]any{
		"roles": roles,
		"orgs":  orgs,
	}
	common.OkWithData(data, c)
}

// SearchUsers 搜索
func (a *UserApi) SearchUsers(c *gin.Context) {
	var data system3.SearchUser
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err, list, total := userService.Search(data)
	if err != nil {
		common.FailWhitStatus(utils2.FindUserError, c)
		return
	}
	common.OkWithData(common.PageVO{
		Items: list,
		Total: total,
	}, c)
}

// CompleteInfo 用户详细信息
func (a *UserApi) CompleteInfo(c *gin.Context) {
	var data system3.UserId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code == utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	err, user := userService.GetCompleteInfoById(data.Id)
	if err != nil {
		common.FailWhitStatus(utils2.FindUserError, c)
		return
	}
	common.OkWithData(user, c)
}

// DeleteUser 删除用户
func (a *UserApi) DeleteUser(c *gin.Context) {
	var data system3.UserId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils2.ParamsResolveFault, c)
		return
	}
	msg, code := utils2.Validate(&data)
	if code != utils2.ERROR {
		common.FailWithMessage(msg.Error(), c)
		return
	}
	if err := userService.Delete(data.Id); err != nil {
		common.FailWhitStatus(utils2.DeleteUserError, c)
		return
	}
	common.Ok(c)
}
