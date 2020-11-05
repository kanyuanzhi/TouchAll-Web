package routers

import (
	"TouchAll-web/controllers"
	"github.com/gin-gonic/gin"
)

func managerGroupStart(routerGroup *gin.RouterGroup) {
	routerGroup.Static("/static", "./static")

	routerGroup.GET("/equipment", controllers.EquipmentManager)
}