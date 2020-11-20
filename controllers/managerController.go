package controllers

import (
	"TouchAll-Web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EquipmentManager(c *gin.Context) {
	equipments := utils.FindAllEquipments()
	equipmentGroupsMap := utils.GetEquipmentGroupsMap()
	equipmentTypesMap := utils.GetEquipmentTypesMap()
	for i := 0; i < len(equipments); i++ {
		equipments[i].EquipmentGroupName, _ = equipmentGroupsMap[equipments[i].EquipmentGroup]
		equipments[i].EquipmentTypeName, _ = equipmentTypesMap[equipments[i].EquipmentType]
	}

	if equipments == nil {
		c.HTML(http.StatusOK, "manager/equipment.html", gin.H{
			"title": "设备管理",
		})
	} else {
		c.HTML(http.StatusOK, "manager/equipment.html", gin.H{
			"title":      "设备管理",
			"equipments": equipments,
		})
	}
}

func SensorManager(c *gin.Context) {
	sensors := utils.FindAllSensors()
	if sensors == nil {
		c.HTML(http.StatusOK, "manager/sensor.html", gin.H{
			"title": "传感器管理",
		})
	} else {
		c.HTML(http.StatusOK, "manager/sensor.html", gin.H{
			"title":   "传感器管理",
			"sensors": sensors,
		})
	}
}
