package routers

import (
	"TouchAll-web/controllers"
	"TouchAll-web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerGroupStart(routerGroup *gin.RouterGroup) {
	routerGroup.Static("/static", "./static")

	routerGroup.GET("/equipment", func(c *gin.Context) {
		equipmentTypes := utils.FindEquipmentTypes()
		equipmentTypesMap := make(map[int]string)
		for _, equipmentType := range equipmentTypes {
			equipmentTypesMap[equipmentType.TypeID] = equipmentType.TypeName
		}

		c.HTML(http.StatusOK, "register/equipment.html", gin.H{
			"title":          "设备注册",
			"equipmentTypes": equipmentTypes,
		})
	})
	routerGroup.POST("/equipment/form", controllers.EquipmentRegister)

	routerGroup.GET("/equipmentType", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register/equipmentType.html", gin.H{
			"title": "设备类型注册",
		})
	})
	routerGroup.POST("/equipment/equipmentTypeForm", controllers.EquipmentTypeRegister)
}
