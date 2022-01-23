package initServe

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"permissions/global"
	"permissions/middleware"
	"time"
)

func RunWindowServer() {
	appConfig := global.System.App
	router := gin.New()
	router.Use(middleware.LogToFile()).Use(gin.Recovery())

	// 静态资源配置
	//router.StaticFS()

	// 实例化路由

	server := &http.Server{
		Addr:           appConfig.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe().Error()
}
