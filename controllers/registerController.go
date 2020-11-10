package controllers

import (
	"TouchAll-Web/models"
	"TouchAll-Web/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"strings"
)

var success bool
var message string

func EquipmentRegister(c *gin.Context) {
	var equipment models.Equipment
	c.ShouldBindWith(&equipment, binding.Form)
	equipment.NetworkMac1 = strings.Replace(equipment.NetworkMac1, " ", "", -1) // 去掉空格
	equipment.NetworkMac2 = strings.Replace(equipment.NetworkMac2, " ", "", -1)
	equipment.NetworkCard1 = strings.Replace(equipment.NetworkCard1, " ", "", -1)
	equipment.NetworkCard2 = strings.Replace(equipment.NetworkCard2, " ", "", -1)
	if equipment.EquipmentID == 0 {
		success = false
		message = "设备ID不能为空！"
	} else if equipment.NetworkMac1 == "" {
		success = false
		message = "设备mac地址1不能为空！"
	} else if utils.IsEquipmentIDExisted(equipment.EquipmentID) {
		success = false
		message = "设备ID重复，请使用其他ID！"
	} else if utils.IsEquipmentNetworkMacExisted(equipment.NetworkMac1, equipment.NetworkMac2) {
		success = false
		message = "该设备mac地址1或2已注册，请转至设备管理页面查看！"
	} else if utils.InsertEquipment(equipment) {
		success = false
		message = "设备注册成功！"
	} else {
		success = false
		message = "设备注册失败！"
	}
	c.JSON(200, gin.H{
		"success": success,
		"message": message,
	})
}

func EquipmentTypeRegister(c *gin.Context) {
	var equipmentType models.EquipmentType
	c.ShouldBindWith(&equipmentType, binding.Form)
	equipmentType.TypeName = strings.Replace(equipmentType.TypeName, " ", "", -1) // 去掉空格
	// equipmentType.TypeID = 0 设备类型编号为空表示"未分类型"
	if equipmentType.TypeName == "" {
		success = false
		message = "设备类型名称不能为空！"
	} else if utils.IsEquipmentTypeExisted(equipmentType.TypeID, equipmentType.TypeName) {
		success = false
		message = "该设备类型编号或名称已注册，请转至设备类型管理页面查看！"
	} else if utils.InsertEquipmentType(equipmentType) {
		success = true
		message = "设备类型注册成功！"
	} else {
		success = false
		message = "设备类型注册失败！"
	}
	c.JSON(200, gin.H{
		"success": success,
		"message": message,
	})
}

func EquipmentGroupRegister(c *gin.Context) {
	var equipmentGroup models.EquipmentGroup
	c.ShouldBindWith(&equipmentGroup, binding.Form)
	equipmentGroup.GroupName = strings.Replace(equipmentGroup.GroupName, " ", "", -1) // 去掉空格
	// equipmentGroup.GroupID = 0 设备组编号为空表示"未分组"
	if equipmentGroup.GroupName == "" {
		success = false
		message = "设备组名称不能为空！"
	} else if utils.IsEquipmentGroupExisted(equipmentGroup.GroupID, equipmentGroup.GroupName) {
		success = false
		message = "该设备组编号或名称已注册，请转至设备组管理页面查看！"
	} else if utils.InsertEquipmentGroup(equipmentGroup) {
		success = true
		message = "设备组注册成功！"
	} else {
		success = false
		message = "设备组注册失败！"
	}
	c.JSON(200, gin.H{
		"success": success,
		"message": message,
	})
}

func CameraRegister(c *gin.Context) {
	var camera models.Camera
	c.ShouldBindWith(&camera, binding.Form)
	camera.CameraName = strings.Replace(camera.CameraName, " ", "", -1) // 去掉空格
	camera.CameraHost = strings.Replace(camera.CameraHost, " ", "", -1)
	if camera.CameraID == 0 {
		success = false
		message = "监控摄像机ID不能为空或0！"
	} else if camera.CameraName == "" {
		success = false
		message = "监控摄像机名称不能为空！"
	} else if camera.CameraHost == "" {
		success = false
		message = "监控摄像机地址不能为空！"
	} else if utils.IsCameraExisted(camera.CameraID, camera.CameraName) {
		success = false
		message = "该监控摄像机编号或名称已注册，请转至监控摄像机管理页面查看！"
	} else if utils.InsertCamera(camera) {
		success = true
		message = "监控摄像机注册成功！"
	} else {
		success = false
		message = "监控摄像机注册失败！"
	}
	c.JSON(200, gin.H{
		"success": success,
		"message": message,
	})
}
