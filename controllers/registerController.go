package controllers

import (
	"TouchAll-web/dbDrivers"
	"TouchAll-web/models"
	"TouchAll-web/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func RegisterEquipment(c *gin.Context) {
	var registrationForm models.EquipmentRegistrationForm
	c.ShouldBindWith(&registrationForm, binding.Form)
	if registrationForm.EquipmentID == 0 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "设备ID不能为空！",
		})
	} else if registrationForm.NetworkMac1 == "" && registrationForm.NetworkMac2 == "" {
		c.JSON(200, gin.H{
			"success": false,
			"message": "设备网卡不能为空！",
		})
	} else if utils.IsEquipmentIDExisted(&registrationForm, dbDrivers.MysqlDB) {
		c.JSON(200, gin.H{
			"success": false,
			"message": "设备ID重复，请使用其他ID！",
		})
	} else {
		success, registeredEquipment := utils.InsertEquipmentRegistration(&registrationForm, dbDrivers.MysqlDB)
		if success {
			if registeredEquipment != nil {
				message := fmt.Sprintf("设备已经存在，%i ，如需修改请转至修改界面！", registrationForm.EquipmentID)
				c.JSON(200, gin.H{
					"success": false,
					"message": message,
				})
			} else {
				c.JSON(200, gin.H{
					"success": true,
					"message": "设备注册成功！",
				})
			}
		} else {
			c.JSON(200, gin.H{
				"success": false,
				"message": "设备注册失败！ ",
			})
		}
	}
}
