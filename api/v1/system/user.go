package system

import (
	"github.com/gin-gonic/gin"
	"permissions/model"
	response "permissions/model/common"
	"permissions/model/system"
	"permissions/utils"
)

type userApi struct{}

// CreateUser 创建用户
func (a *userApi) CreateUser(c *gin.Context) {
	var data system.User
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&data)
	if code == utils.ERROR {
		response.FailWithMessage(msg.Error(), c)
	}
	var roles []system.SysRole
	var orgs []system.SysOrganize
	for _, roleId := range data.SysRoleIds {
		roles = append(roles, system.SysRole{BaseID: model.BaseID{ID: roleId}})
	}
	for _, orgId := range data.SysOrganizeIds {
		orgs = append(orgs, system.SysOrganize{BaseID: model.BaseID{ID: orgId}})
	}
	if err := userService.Register(&system.SysUser{
		Username:     data.Username,
		LoginName:    data.LoginName,
		Password:     data.Password,
		SysRoles:     roles,
		SysOrganizes: orgs,
	}); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.Ok(c)
}

// UpdateUserBaseInfo 更新基本用户
func (a *userApi) UpdateUserBaseInfo(c *gin.Context) {
	var data system.UserBaseInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	if err := userService.UpdateUserInfo(system.SysUser{
		BaseUUID:  model.BaseUUID{ID: data.Id},
		Username:  data.Username,
		Password:  data.Password,
		LoginName: data.LoginName,
	}); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.Ok(c)
}

// 更新用户权限
func (a *userApi) UpdateUserPerInfo(c *gin.Context) {
	var data system.UserPerInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	if err := userService.SetUserRoleAndOrg(data.Id, data.SysRoleIds, data.SysOrganizeIds); err != nil {
		response.Fail(c)
	}
	response.Ok(c)
}

// Login 登录
func (a *userApi) Login(c *gin.Context) {
	var data system.UserLogin
	if err := c.ShouldBindJSON(&data); err != nil {
		response.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	err, user := userService.GetUserByLoginName(data.LoginName)
	if err != nil {
		response.FailWhitStatus(utils.UsernameAndPasswdError, c)
	}
	if user.Password != data.Password {
		response.FailWhitStatus(utils.UsernameAndPasswdError, c)
	}
	//response.OkWithData(, c)
}

// GetOLUserInfo 当前登录用户完整信息
func (a userApi) GetOLUserInfo() {

}
