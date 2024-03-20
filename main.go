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
	r.GET("/products", controllers.GetProducts)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.Run()
}
