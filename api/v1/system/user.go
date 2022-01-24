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
	var user system.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWhitStatus(utils.ParamsResolveFault, c)
	}
	msg, code := utils.Validate(&user)
	if code == utils.ERROR {
		response.FailWithMessage(msg.Error(), c)
	}
	var roles []system.SysRole
	var orgs []system.SysOrganize
	for _, roleId := range user.SysRoleIds {
		roles = append(roles, system.SysRole{BaseID: model.BaseID{ID: roleId}})
	}
	for _, orgId := range user.SysOrganizeIds {
		orgs = append(orgs, system.SysOrganize{BaseID: model.BaseID{ID: orgId}})
	}
	err := userService.Register(&system.SysUser{
		Username:     user.Username,
		LoginName:    user.LoginName,
		Password:     user.Password,
		SysRoles:     roles,
		SysOrganizes: orgs,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.Ok(c)
}

// 更新基本用户
func (a *userApi) UpdateBaseInfo(c *gin.Context) {

}

// 更新用户权限
