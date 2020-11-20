package routers

import (
	"TouchAll-Web/controllers"
	"TouchAll-Web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// response for "/register"
func registerGroupStart(routerGroup *gin.RouterGroup) {
	routerGroup.Static("/static", "./static")

	// 注册设备
	routerGroup.GET("/equipment", func(c *gin.Context) {
		equipmentTypes := utils.FindEquipmentTypes()
		equipmentGroups := utils.FindEquipmentGroups()
		c.HTML(http.StatusOK, "register/equipment.html", gin.H{
			"title":           "设备注册",
			"equipmentTypes":  equipmentTypes,
			"equipmentGroups": equipmentGroups,
		})
	})
	routerGroup.POST("/equipment", controllers.EquipmentRegister)

	// 注册设备名称
	routerGroup.GET("/equipmentType", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register/equipmentType.html", gin.H{
			"title": "设备类型注册",
		})
	})
	routerGroup.POST("/equipmentType", controllers.EquipmentTypeRegister)

	// 注册设备组
	routerGroup.GET("/equipmentGroup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register/equipmentGroup.html", gin.H{
			"title": "设备组注册",
		})
	})
	routerGroup.POST("/equipmentGroup", controllers.EquipmentGroupRegister)

	// 注册监控摄像机
	routerGroup.GET("/camera", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register/camera.html", gin.H{
			"title": "监控摄像机注册",
		})
	})
	routerGroup.POST("/camera", controllers.CameraRegister)

	// 注册监控摄像机
	routerGroup.GET("/sensor", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register/sensor.html", gin.H{
			"title": "传感器注册",
		})
	})
	routerGroup.POST("/sensor", controllers.SensorRegister)
}
