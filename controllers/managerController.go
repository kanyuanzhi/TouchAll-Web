package controllers

import (
	"TouchAll-web/utils"
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
		c.HTML(http.StatusOK, "manager/equipment.html", gin.H{})
	} else {
		c.HTML(http.StatusOK, "manager/equipment.html", gin.H{
			"title":      "设备管理",
			"equipments": equipments,
		})
	}
}
