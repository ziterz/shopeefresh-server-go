package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ziterz/shopeefresh-server-go/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"products": []string{},
		})
	})
	r.Run()
}
