package controllers

import (
	"TouchAll-web/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EquipmentManager(c *gin.Context) {
	equipments := utils.FindAllEquipments()
	if equipments == nil{
		c.HTML(http.StatusOK, "manager/equipment.html", gin.H{})
	}else {
		c.HTML(http.StatusOK, "manager/equipment.html", gin.H{
			"title":"设备管理",
			"equipments":equipments,
		})
	}
}
