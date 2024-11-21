// controllers/orders.go

package controllers

import (
	"net/http"

	"github.com/Acstrayer/TESCSE-Ecom/api/models"
	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	ID uint   `json:"id" binding:"required"`
	Items []OrderItem `json:"items" binding:"required"`
      Name string `json:"name" binding:"required"`
      Email string `json:"email" binding:"required"`
      Address1 string `json:"address_one" binding:"required"`
      Address2 string `json:"address_two" binding:"required"`
      City string `json:"city" binding:"required"`
      State string `json:"state" binding:"required"`
      Zip string `json:"zip" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	// Validate input
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Post order
	order := models.Order{Items: input.Items}
	models.DB.Create(&order)

	c.JSON(http.StatusOK, gin.H{"data": order})
}
