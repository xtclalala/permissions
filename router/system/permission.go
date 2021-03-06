package system

import (
	"github.com/gin-gonic/gin"
	"permissions/api/v1"
	//"permissions/middleware"
)

type PermissionRouter struct{}

func (r *PermissionRouter) InitPerRouter(router *gin.RouterGroup) {
	userRouter := router.Group("permission") //.Use(middleware.LogToFile())
	var perApi = v1.ApiGroupApp.SysApiGroup.PermissionApi
	{
		userRouter.POST("per", perApi.Register)
		userRouter.PUT("per", perApi.UpdatePerBaseInfo)
		userRouter.GET("perAll", perApi.PermissionAllByMenuId)
		userRouter.DELETE("per", perApi.DeletePermission)
		userRouter.GET("per", perApi.SearchPermission)
	}
}
