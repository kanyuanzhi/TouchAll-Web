package routers

import (
	"TouchAll-Web/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func displayGroupStart(routerGroup *gin.RouterGroup) {
	routerGroup.Static("/static", "./static")

	routerGroup.GET("/equipment", controllers.SingleEquipmentDisplay)

	routerGroup.GET("/allEquipments", controllers.AllEquipmentsDisplay)

	routerGroup.GET("/dataCenter", func(c *gin.Context) {
		c.HTML(http.StatusOK, "display/dataCenterWsConnectionStatus.html", gin.H{
			"title": "数据中心websocket连接状态",
		})
	})

	routerGroup.GET("/camera", controllers.CameraDisplay)
}
