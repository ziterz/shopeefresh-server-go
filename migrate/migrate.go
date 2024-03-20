package main

import (
	"github.com/ziterz/shopeefresh-server-go/initializers"
	"github.com/ziterz/shopeefresh-server-go/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Product{})
}
