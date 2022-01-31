package system

import (
	"github.com/gin-gonic/gin"
	"permissions/global"
	"permissions/model/common"
	"permissions/model/system"
	system2 "permissions/services/system"
	"permissions/utils"
)

type UserApi struct{}

// CreateUser 创建用户
func (a *UserApi) CreateUser(c *gin.Context) {
	var data system.User
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
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
		Password:     data.Password,
		SysRoles:     roles,
		SysOrganizes: orgs,
	}); err != nil {
		common.FailWithMessage(err.Error(), c)
	}
	common.Ok(c)
}

// UpdateUserBaseInfo 更新基本用户
func (a *UserApi) UpdateUserBaseInfo(c *gin.Context) {
	var data system.UserBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := userService.UpdateUserInfo(system.SysUser{
		BaseUUID:  system.BaseUUID{ID: data.Id},
		Username:  data.Username,
		Password:  data.Password,
		LoginName: data.LoginName,
	}); err != nil {
		common.FailWithMessage(err.Error(), c)
	}
	common.Ok(c)
}

// 更新用户权限
func (a *UserApi) UpdateUserPerInfo(c *gin.Context) {
	var data system.UserPerInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := userService.SetUserRoleAndOrg(data.Id, data.SysRoleIds, data.SysOrganizeIds); err != nil {
		common.Fail(c)
	}
	common.Ok(c)
}

// Login 登录
func (a *UserApi) Login(c *gin.Context) {
	var data system.UserLogin
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, user := userService.GetUserByLoginName(data.LoginName)
	if err != nil {
		common.FailWhitStatus(utils.UsernameAndPasswdError, c)
	}
	if user.Password != data.Password {
		common.FailWhitStatus(utils.UsernameAndPasswdError, c)
	}
	j := system2.NewJWT()
	claim := j.CreateClaim(&user)
	token, err := j.CreateJwt(&claim)
	if err != nil {
		common.FailWhitStatus(utils.TokenCreateFailed, c)
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
	}
	ok, orgs := organizeService.GetOrgByUserId(claim.Id)
	if ok != nil {
		common.FailWhitStatus(utils.FindOrgError, c)
	}
	roles := make([]system.SysRole, 10)
	for _, org := range orgs {
		ok, sysRoles := roleService.GetRoleByOrgId(org.ID)
		if ok != nil {
			common.FailWhitStatus(utils.FindRoleError, c)
		}
		roles = append(roles, sysRoles...)
	}
	pers := make([]system.SysPermission, 200)
	for _, role := range roles {
		ok, SysPermissions := permissionService.GetPerByRoleId(role.ID)
		if ok != nil {
			common.FailWhitStatus(utils.FindPermissionError, c)
		}
		pers = append(pers, SysPermissions...)
	}
	menus := make([]system.SysMenu, 50)
	for _, role := range roles {
		ok, sysMenus := menuService.GetMenuByRoleId(role.ID)
		if ok != nil {
			common.FailWhitStatus(utils.FindMenuError, c)
		}
		menus = append(menus, sysMenus...)
	}
	data := map[string]interface{}{
		"menus": menus,
		"pers":  pers,
		"roles": roles,
		"orgs":  orgs,
	}
	common.OkWithData(data, c)
}

// SearchUsers 搜索
func (a *UserApi) SearchUsers(c *gin.Context) {
	var data system.SearchUser
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, list, total := userService.Search(data)
	if err != nil {
		common.FailWhitStatus(utils.FindUserError, c)
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
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	err, user := userService.GetCompleteInfoById(data.Id)
	if err != nil {
		common.FailWhitStatus(utils.FindUserError, c)
	}
	common.OkWithData(user, c)
}

// DeleteUser 删除用户
func (a *UserApi) DeleteUser(c *gin.Context) {
	var data system.UserId
	if err := c.ShouldBindJSON(&data); err != nil {
		common.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code != utils.ERROR {
		common.FailWithMessage(msg.Error(), c)
	}
	if err := userService.Delete(data.Id); err != nil {
		common.FailWhitStatus(utils.DeleteUserError, c)
	}
	common.Ok(c)
}
