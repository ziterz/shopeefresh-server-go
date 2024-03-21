package main

import (
	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products", controllers.GetProducts)
	r.GET("/products/:id", controllers.GetProductByID)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.Run()
}
