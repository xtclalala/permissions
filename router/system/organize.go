package system

import (
	"github.com/gin-gonic/gin"
	"permissions/api/v1"
)

type OrganizeRouter struct{}

func (r *OrganizeRouter) InitOrganizeRouter(router *gin.RouterGroup) {
	userRouter := router.Group("organize") //.Use(middleware.LogToFile())
	var organizeApi = v1.ApiGroupApp.SysApiGroup.OrganizeApi
	{
		userRouter.POST("register", organizeApi.Register)
		userRouter.PUT("organize", organizeApi.UpdateOrgBaseInfo)
		userRouter.POST("organize", organizeApi.SearchOrganize)
		userRouter.DELETE("organize", organizeApi.DeleteOrganize)
	}
}
