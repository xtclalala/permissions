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
		userRouter.POST("register", userApi.CreateUser)
		userRouter.POST("baseInfo", userApi.UpdateUserBaseInfo)
		userRouter.POST("perInfo", userApi.UpdateUserPerInfo)
		userRouter.POST("login", userApi.Login)
		userRouter.GET("routerAndRole", userApi.GetUserRouterAndRoles)
		userRouter.POST("SearchUsers", userApi.SearchUsers)
		userRouter.GET("user", userApi.CompleteInfo)
		userRouter.DELETE("user", userApi.DeleteUser)

	}
}
