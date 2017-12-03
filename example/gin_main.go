package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main3() {
	// Log相关设置
	// Disable Console Color, you don't need console when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Use the following code if you need to write the logs to file and console at the same time.
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	// 分组或分版本管理
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name/", PingHandler)
		v1.POST("/post", PostHandler)
		v1.POST("form_post", FormPostHandler)
		v1.POST("upload", UploadFileHandler)
		v1.POST("uploadmutifile", UploadMutiFileHandler)
	}

	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}

func PingHandler(c *gin.Context) {
	// 这些字段都是urlencode的
	// 这个是解析path的
	pathName := c.Param("name")
	// Query 是解析参数的
	param := c.DefaultQuery("name", "defaultname")
	c.String(http.StatusOK, "Hello String")
	c.JSON(200, gin.H{
		"message": "pong",
		"name":    pathName,
		"param":   param,
	})
}

func PostHandler(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	// x-www-rorm-urlencoded 或者 form-data都可以，如果不成功，请检查你的Header里Content-Type(如果写了某种类型，请使用对应类型去发送)
	// 特别是在postman里会自动带一个Content-Type的情况。
	name := c.PostForm("name")
	message := c.PostForm("message")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"page":    page,
		"name":    name,
		"message": message,
	})
}

type s_data struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func FormPostHandler(c *gin.Context) {
	var p s_data
	c.BindJSON(&p)
	fmt.Println("p", p)
	fmt.Println()
	fmt.Println()
	c.JSON(http.StatusOK, gin.H{
		"name":    p.Name,
		"message": p.Message,
	})
}

func UploadFileHandler(c *gin.Context) {
	file, _ := c.FormFile("file")

	// dest需要包含文件名，不能是路径
	c.SaveUploadedFile(file, fmt.Sprintf("C:/%s", file.Filename))

	c.String(http.StatusOK, fmt.Sprintf("%s uploaded!", file.Filename))
}

func UploadMutiFileHandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		c.SaveUploadedFile(file, fmt.Sprintf("C:/%s", file.Filename))
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
