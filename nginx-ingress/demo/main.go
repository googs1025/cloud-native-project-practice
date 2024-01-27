package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "root"})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "user"})
	})

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "test"})
	})

	r.Run(":8080")
}
