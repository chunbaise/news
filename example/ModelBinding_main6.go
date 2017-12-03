package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Can Bind form & json
type Login struct {
	User     string `form:"user" json:"User" binding:"required"`         //
	Password string `form:"password" json:"Password" binding:"required"` //
}

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main6() {
	router := gin.Default()

	router.POST("/loginjson", LoginJsonHandler)
	router.POST("/loginform", LoginFormHandler)
	router.Any("/testing", StatrPageHandler)

	router.Run(":8082")
}

func StatrPageHandler(c *gin.Context) {
	var person Person
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "Success")
}

func LoginJsonHandler(c *gin.Context) {
	var json Login

	if err := c.ShouldBindJSON(&json); err == nil {
		if json.User == "manu" && json.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
}

func LoginFormHandler(c *gin.Context) {
	var form Login
	// 这里会根据Content-type来作使用哪个类型来Binding
	if err := c.ShouldBind(&form); err == nil {
		if form.User == "manu" && form.Password == "123" {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	}
}
