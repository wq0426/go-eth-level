package main

import (
	"fmt"
	"rcc/level2/gin/middlewire"
	model "rcc/level2/gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/getb", model.GetDataB)
		v1.GET("/getc", model.GetDataC)
		v1.GET("/getd", model.GetDataD)
	}
	r.Use(middlewire.Logger())
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
