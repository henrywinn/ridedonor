package main

import (
    "gopkg.in/gin-gonic/gin.v1"
    _ "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" 
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ponggggg",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
