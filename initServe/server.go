package initServe

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"permissions/global"
	"permissions/middleware"
	Router "permissions/router"
	"time"
)

func RunWindowServer() {
	appConfig := global.System.App
	router := gin.New()
	router.Use(middleware.LogToFile()).Use(gin.Recovery())
	routerGroup := router.Group("")
	// 静态资源配置
	//router.StaticFS()
	// 实例化路由
	sysRouter := Router.AppRouter.System
	sysRouter.InitUserRouter(routerGroup)     // 用户路由
	sysRouter.InitPerRouter(routerGroup)      // 按钮路由
	sysRouter.InitMenuRouter(routerGroup)     // 菜单路由
	sysRouter.InitRoleRouter(routerGroup)     // 角色路由
	sysRouter.InitOrganizeRouter(routerGroup) // 组织路由

	server := &http.Server{
		Addr:           appConfig.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe().Error()
}
