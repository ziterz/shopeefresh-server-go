package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ziterz/shopeefresh-server-go/controllers"
	"github.com/ziterz/shopeefresh-server-go/initializers"
	"github.com/ziterz/shopeefresh-server-go/middlewares"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/products", middlewares.Authentication, controllers.CreateProduct)
	r.GET("/products", middlewares.Authentication, controllers.GetProducts)
	r.PUT("/products/:id", middlewares.Authentication, controllers.UpdateProduct)
	r.DELETE("/products/:id", middlewares.Authentication, controllers.DeleteProduct)
	r.Run()
}
