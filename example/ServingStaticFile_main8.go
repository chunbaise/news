package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main8() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("C:/"))
	router.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
