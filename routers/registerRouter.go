package routers

import (
	"TouchAll-web/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerGroupStart(router *gin.RouterGroup) {
	router.Static("/static", "./static")
	router.GET("/equipment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register/equipment.html", gin.H{})
	})
	router.POST("/equipment/form", controllers.RegisterEquipment)
}
