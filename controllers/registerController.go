package controllers

import (
	"TouchAll-web/models"
	"TouchAll-web/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func EquipmentRegister(c *gin.Context) {
	var equipment models.Equipment
	c.ShouldBindWith(&equipment, binding.Form)
	if equipment.EquipmentID == 0 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "设备ID不能为空！",
		})
	} else if equipment.NetworkMac1 == "" && equipment.NetworkMac2 == "" {
		c.JSON(200, gin.H{
			"success": false,
			"message": "设备网卡不能为空！",
		})
	} else if utils.IsEquipmentIDExisted(equipment.EquipmentID) {
		c.JSON(200, gin.H{
			"success": false,
			"message": "设备ID重复，请使用其他ID！",
		})
	} else if utils.IsEquipmentNetworkMacExisted(equipment.NetworkMac1, equipment.NetworkMac2) {
		c.JSON(200, gin.H{
			"success": false,
			"message": "该设备mac地址已注册，请转至设备管理页面查看！",
		})
	} else if utils.InsertEquipment(equipment) {
		c.JSON(200, gin.H{
			"success": true,
			"message": "设备注册成功！",
		})
	} else {
		c.JSON(200, gin.H{
			"success": false,
			"message": "设备注册失败！ ",
		})
	}
}
