package main

import (
	"github.com/gin-gonic/gin"
)

func main2() {
	r := gin.Default()
	r.GET("/ping/:name/*action", PingHandler)
	r.POST("/post/:name", PostHandler)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}

func PingHandler1(c *gin.Context) {
	pathName := c.Param("name")
	action := c.Param("action")
	action2 := c.Params.ByName("action")
	param := c.DefaultQuery("name", "defaultname")
	lastname := c.Query("lastname")
	c.JSON(200, gin.H{
		"message":  "pong",
		"name":     pathName,
		"action":   action,
		"action2":  action2,
		"param":    param,
		"lastname": lastname,
	})
}

func PostHandler1(c *gin.Context) {
	pathName := c.Param("name")
	param := c.DefaultQuery("name", "defaultname")
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
		"name":    pathName,
		"param":   param,
	})
}
