package system

import (
	"github.com/gin-gonic/gin"
	"permissions/api/v1"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user") //.Use(middleware.LogToFile())
	var userApi = v1.ApiGroupApp.SysApiGroup.UserApi
	{
		userRouter.POST("user", userApi.CreateUser)
		userRouter.PUT("user", userApi.UpdateUserBaseInfo)
		userRouter.PUT("per", userApi.UpdateUserPerInfo)
		userRouter.GET("routerAndRole", userApi.GetUserRouterAndRoles)
		userRouter.GET("user", userApi.SearchUsers)
		userRouter.GET("completeInfo", userApi.CompleteInfo)
		userRouter.DELETE("user", userApi.DeleteUser)

	}
}
