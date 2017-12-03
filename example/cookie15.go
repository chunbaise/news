package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/getcookie", PingHandler3)
	r.POST("/postcookie", PostHandler3)
	r.Run(":8082") // listen and serve on 0.0.0.0:8080
}

func PingHandler3(c *gin.Context) {
	COOKIE_MAX_MAX_AGE := time.Hour * 24 / time.Second // 单位：秒。
	maxAge := int(COOKIE_MAX_MAX_AGE)
	uid := "10"

	uid_cookie := &http.Cookie{
		Name:     "uid",
		Value:    uid,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   maxAge,
	}

	// cookie是自动带上的，因为cookie设计的目录就是这样的
	cookiefromclient, _ := c.Request.Cookie("uid")
	http.SetCookie(c.Writer, uid_cookie)
	c.JSON(200, gin.H{
		"message":      "pong",
		"getthecookie": cookiefromclient,
	})
}

func PostHandler3(c *gin.Context) {
	uid := "20"

	uid_cookie := &http.Cookie{
		Name:     "uid2",
		Value:    uid,
		Path:     "/",
		HttpOnly: false,
	}

	http.SetCookie(c.Writer, uid_cookie)
	c.JSON(200, gin.H{
		"status": "posted",
	})
}
