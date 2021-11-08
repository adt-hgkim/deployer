package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	relativePath := "./public"
	pattern := "./public/**/index.html"

	args := os.Args[1:]
	if len(args) > 0 {
		gin.SetMode(gin.ReleaseMode)
		relativePath = args[0] + "/public"
		pattern = args[0] + "/public/**/index.html"
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
