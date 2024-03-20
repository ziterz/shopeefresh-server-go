package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ziterz/shopeefresh-server-go/controllers"
	"github.com/ziterz/shopeefresh-server-go/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/products", controllers.CreateProduct)
	r.Run()
}
