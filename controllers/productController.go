package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ziterz/shopeefresh-server-go/initializers"
	"github.com/ziterz/shopeefresh-server-go/models"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name  string
		Image string
		Price int
	}
	c.Bind(&body)

	// Create product
	product := models.Product{
		Name:  body.Name,
		Image: body.Image,
		Price: body.Price,
	}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func GetProducts(c *gin.Context) {

}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
