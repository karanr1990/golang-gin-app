package main

import (
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run()
}

