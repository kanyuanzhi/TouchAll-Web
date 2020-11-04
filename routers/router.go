package routers

import "github.com/gin-gonic/gin"

type Router struct {
	r *gin.Engine
}

func NewRouter() *Router {
	return &Router{
		r: gin.Default(),
	}
}

func (router *Router) Start() {
	router.r.LoadHTMLGlob("templates/**/*")
	router.r.Static("/static", "./static")
	//router.r.Static("/register", "./static")
	registerGroup := router.r.Group("/register")
	registerGroupStart(registerGroup)

	router.r.Run(":8081")
}
