package initServe

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"permissions/global"
	"permissions/middleware"
	"time"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunWindowServer() {
	appConfig := global.System.App
	router := gin.New()

	// 静态资源配置
	//router.StaticFS()
	// 配置路由
	// 公共路由
	publicGroup := router.Group("")
	//var userApi = v1.ApiGroupApp.SysApiGroup.UserApi
	{
		publicGroup.GET("/ws/:documentId", func(c *gin.Context) {
			docId := c.Param("documentId")
			log.Print("docId:", docId)
			ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
			if err != nil {
				log.Print("upgrade:", err)
				return
			}
			defer ws.Close()
			for {
				mt, message, err := ws.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					break
				}
				log.Printf("recv: %s", message)
				err = ws.WriteMessage(mt, message)
				if err != nil {
					log.Println("write:", err)
					break
				}
			}
		})
		//publicGroup.POST("login", userApi.Login)
	}
	// 私有路由
	//router.Use(middleware.Auth())
	router.Use(middleware.Cors())
	router.Use(middleware.LogToFile()).Use(gin.Recovery())
	//privateGroup := router.Group("")
	//sysRouter := Router.AppRouter.System
	//sysRouter.InitUserRouter(privateGroup)     // 用户路由
	//sysRouter.InitPerRouter(privateGroup)      // 按钮路由
	//sysRouter.InitMenuRouter(privateGroup)     // 菜单路由
	//sysRouter.InitRoleRouter(privateGroup)     // 角色路由
	//sysRouter.InitOrganizeRouter(privateGroup) // 组织路由

	server := &http.Server{
		Addr:           appConfig.Port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = server.ListenAndServe().Error()
}
