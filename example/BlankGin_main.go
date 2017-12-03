package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main4() {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// 单个接口
	// Per route middleware, you can add as many as you desire.
	r.GET("/benchmark", MyBenchLogger(), benchEndPoint)

	// 管理一组接口
	authorized := r.Group("/authorized")
	authorized.Use(AuthRequired())
	{
		authorized.POST("login", loginEndpoint)
		authorized.POST("getinfo", getinfoEndpoint)

		//  这里还可以有嵌套的分组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}
	r.Run(":8082")
}

func analyticsEndpoint(c *gin.Context) {
	// 在子分组中
	fmt.Println("In nested group")
}
func loginEndpoint(c *gin.Context) {
	// Do login
	fmt.Println("Do login")
}

func getinfoEndpoint(c *gin.Context) {
	// Return info
	fmt.Println("Return Info")
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Do Something Auth Check
		fmt.Printf("Do Auth Check\n")
	}
}
func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func benchEndPoint(c *gin.Context) {
	c.String(http.StatusOK, "benchEndPoint")
}
