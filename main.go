package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	relativePath := "./public"
	pattern := "./public/**/index.html"

	argsWithoutProg := os.Args[1:]
	for i := 0; i < len(argsWithoutProg); i++ {
		if argsWithoutProg[i] == "release" {
			gin.SetMode(gin.ReleaseMode)
			relativePath = "/home/pi/deployer/public"
			pattern = "/home/pi/deployer/public/**/index.html"
		}
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "deployer",
		})
	})
	router.Static("/public", relativePath)
	router.LoadHTMLGlob(pattern)

	err := router.Run()
	if err != nil {
		return
	}
}
