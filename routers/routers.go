package routers

import (
	"TouchAll-web/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	r *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		r: gin.Default(),
	}
}

func (router *Router) Start() {
	config := config.NewConfig()
	port := config.GetValue("web_server.port").(string)

	router.r.Static("/static", "./static")
	router.r.LoadHTMLGlob("templates/**/*")

	router.r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", gin.H{})
	})

	// 注册路由组
	registerGroup := router.r.Group("/register")
	registerGroupStart(registerGroup)

	// 管理路由组
	managerGroup := router.r.Group("/manager")
	managerGroupStart(managerGroup)

	// websocket实时更新路由组
	displayGroup := router.r.Group("/display")
	displayGroupStart(displayGroup)

	router.r.Run(":" + port)
}
