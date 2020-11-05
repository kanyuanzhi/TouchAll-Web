package routers

import (
	"TouchAll-web/controllers"
	"github.com/gin-gonic/gin"
)

func displayGroupStart(routerGroup *gin.RouterGroup) {
	routerGroup.Static("/static", "./static")

	routerGroup.GET("/equipment", controllers.EquipmentDisplay)

	routerGroup.GET("/allEquipments", controllers.AllEquipmentsDisplay)
}