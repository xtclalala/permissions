package initServe

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "permissions/api/v1"
	"permissions/global"
	"permissions/middleware"
	Router "permissions/router"
	"time"
)

func RunWindowServer() {
	appConfig := global.System.App
	router := gin.New()

	// 静态资源配置
	//router.StaticFS()
	// 配置路由
	// 公共路由
	publicGroup := router.Group("")
	var userApi = v1.ApiGroupApp.SysApiGroup.UserApi
	{
		publicGroup.POST("login", userApi.Login)
	}
	// 私有路由
	router.Use(middleware.Auth())
	router.Use(middleware.Cors())
	router.Use(middleware.LogToFile()).Use(gin.Recovery())
	privateGroup := router.Group("")
	sysRouter := Router.AppRouter.System
	sysRouter.InitUserRouter(privateGroup)     // 用户路由
	sysRouter.InitPerRouter(privateGroup)      // 按钮路由
	sysRouter.InitMenuRouter(privateGroup)     // 菜单路由
	sysRouter.InitRoleRouter(privateGroup)     // 角色路由
	sysRouter.InitOrganizeRouter(privateGroup) // 组织路由

	server := &http.Server{
		Addr:           appConfig.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe().Error()
}
