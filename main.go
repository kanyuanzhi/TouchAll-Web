package main

import (
	"TouchAll-web/dbDrivers"
	"TouchAll-web/routers"
)

//var useMongodb = config.GetValue("mongodb.use").(bool)
//var useMysql = config.GetValue("mysql.use").(bool)

var mysqlConn = dbDrivers.GetMysqlConn()

func main() {
	router := routers.NewRouter()

	router.Start()

	//config := utils.NewConfig()
	//port := config.GetValue("web_server.port").(string)

	//r := gin.Default()
	//r.LoadHTMLGlob("template/**/*")
	//r.Static("/static", "./static")
	//
	//r.GET("/main", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "video.html", gin.H{})
	//})
	//
	//r.GET("/flv", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "flv.html", gin.H{})
	//})
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.GET("/wsPerson", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "wsPerson.html", gin.H{})
	//})
	//r.GET("/wsPeople", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "wsPeople.html", gin.H{})
	//})
	//
	//r.GET("/wsEquipment", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "wsEquipment.html", gin.H{})
	//})
	//
	//r.GET("/wsVideo", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "wsVideo.html", gin.H{})
	//})
	//
	//r.GET("/registerEquipment", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "equipmentRegistration.html", gin.H{})
	//
	//})
	//
	//r.POST("/postEquipmentRegistrationForm", func(c *gin.Context) {
	//	var registrationForm models.EquipmentRegistrationForm
	//	c.ShouldBindWith(&registrationForm, binding.Form)
	//	if registrationForm.EquipmentID == 0 {
	//		c.JSON(200, gin.H{
	//			"success": false,
	//			"message": "设备ID不能为空！",
	//		})
	//	} else if registrationForm.NetworkMac1 == "" && registrationForm.NetworkMac2 == "" {
	//		c.JSON(200, gin.H{
	//			"success": false,
	//			"message": "设备网卡不能为空！",
	//		})
	//	} else if utils.IsEquipmentIDExisted(&registrationForm, mysqlConn) {
	//		c.JSON(200, gin.H{
	//			"success": false,
	//			"message": "设备ID重复，请使用其他ID！",
	//		})
	//	} else {
	//		success, registeredEquipment := utils.InsertEquipmentRegistration(&registrationForm, mysqlConn)
	//		if success {
	//			if registeredEquipment != nil {
	//				message := fmt.Sprintf("设备已经存在，%i ，如需修改请转至修改界面！", registrationForm.EquipmentID)
	//				c.JSON(200, gin.H{
	//					"success": false,
	//					"message": message,
	//				})
	//			} else {
	//				c.JSON(200, gin.H{
	//					"success": true,
	//					"message": "设备注册成功！",
	//				})
	//			}
	//		} else {
	//			c.JSON(200, gin.H{
	//				"success": false,
	//				"message": "设备注册失败！ ",
	//			})
	//		}
	//	}
	//})
	//
	//r.Run(":" + port)
}
