package system

import (
	"github.com/gin-gonic/gin"
	"permissions/api/v1"
)

type MenuRouter struct{}

func (r *MenuRouter) InitMenuRouter(router *gin.RouterGroup) {
	userRouter := router.Group("menu") //.Use(middleware.LogToFile())
	var menuApi = v1.ApiGroupApp.SysApiGroup.MenuApi
	{
		userRouter.POST("menu", menuApi.Register)
		userRouter.PUT("menu", menuApi.UpdateMenuInfo)
		userRouter.GET("menu", menuApi.SearchMenu)
		userRouter.GET("allMenu", menuApi.AllMenu)
		userRouter.DELETE("menu", menuApi.DeleteMenu)
	}
}
