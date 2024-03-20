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
	c.JSON(http.StatusCreated, gin.H{
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	// Get the product
	products := []models.Product{}
	initializers.DB.Find(&products)

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func GetProductByID(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the product
	product := []models.Product{}
	find := initializers.DB.First(&product, id)
	if find.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not found",
		})
		return

	}

	// Respond with them
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get data off req body
	var body struct {
		Name  string
		Image string
		Price int
	}
	c.Bind(&body)

	// Get the product
	product := models.Product{}
	initializers.DB.First(&product, id)

	// Update the product
	result := initializers.DB.Model(&product).Updates(models.Product{
		Name:  body.Name,
		Image: body.Image,
		Price: body.Price,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update the product",
		})
		return
	}

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func DeleteProduct(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	// Get the product
	product := models.Product{}
	find := initializers.DB.First(&product, id)
	if find.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Delete the product
	result := initializers.DB.Delete(&product, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete the product",
		})
		return
	}

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted",
		"product": product,
	})
}
