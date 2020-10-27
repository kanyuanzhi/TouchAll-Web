package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.Static("/static","./static")

	r.GET("/main", func(c *gin.Context) {
		c.HTML(http.StatusOK, "video.html", gin.H{})
	})

	r.GET("/flv", func(c *gin.Context) {
		c.HTML(http.StatusOK, "flv.html", gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/wsPerson", func(c *gin.Context) {
		c.HTML(http.StatusOK, "wsPerson.html", gin.H{})
	})
	r.GET("/wsPeople", func(c *gin.Context) {
		c.HTML(http.StatusOK, "wsPeople.html", gin.H{})
	})

	r.GET("/wsEquipment", func(c *gin.Context) {
		c.HTML(http.StatusOK, "wsEquipment.html", gin.H{})
	})

	r.GET("/wsVideo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "wsVideo.html", gin.H{})
	})
	r.Run(":8081")
}
