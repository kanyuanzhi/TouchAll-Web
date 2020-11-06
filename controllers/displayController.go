package controllers

import (
	"TouchAll-web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func EquipmentDisplay(c *gin.Context) {
	equipmentIDStr := c.DefaultQuery("equipmentID", "0")
	equipmentID, _ := strconv.Atoi(equipmentIDStr)
	if equipment, has := utils.FindEquipmentByID(equipmentID); has {
		c.HTML(http.StatusOK, "display/equipment.html", gin.H{
			"title":         "单个设备状态展示",
			"equipment":     equipment,
			"equipment_ids": []int{equipmentID},
		})
	} else {
		c.String(http.StatusOK, "没有id为%s的设备", equipment.EquipmentID)
	}
}

func AllEquipmentsDisplay(c *gin.Context) {
	equipments := utils.FindAllEquipmentIDs()
	equipmentIDs := make([]int, 0)
	for _, equipment := range equipments {
		equipmentIDs = append(equipmentIDs, equipment.EquipmentID)
	}
	c.HTML(http.StatusOK, "display/allEquipments.html", gin.H{
		"title":         "所有设备状态展示",
		"equipments":    equipments,
		"equipment_ids": equipmentIDs,
	})
}
