package system

import (
	"github.com/gin-gonic/gin"
	v1 "permissions/api/v1"
	//"permissions/middleware"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(router *gin.RouterGroup) {
	userRouter := router.Group("role") //.Use(middleware.LogToFile())
	var roleApi = v1.ApiGroupApp.SysApiGroup.RoleApi
	{
		userRouter.POST("register", roleApi.CreateRole)
		userRouter.POST("roleBaseInfo", roleApi.UpdateBaseRole)
		userRouter.POST("role", roleApi.UpdateRoleMenus)
		userRouter.POST("copyRole", roleApi.CopyRole)
		userRouter.GET("roleCompleteInfo", roleApi.CompleteRole)
		userRouter.DELETE("role", roleApi.DeleteRole)
		userRouter.GET("role", roleApi.SearchRole)
		userRouter.GET("roleByOrg", roleApi.RoleAllByOrg)
	}
}
