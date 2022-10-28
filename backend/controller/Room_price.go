package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /Room_types

func CreateRoom_price(c *gin.Context) {

	var room_prices entity.Room_price

	if err := c.ShouldBindJSON(&room_prices); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&room_prices).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})

}

// GET /Room_type/:id

func GetRoom_price(c *gin.Context) {

	var room_prices entity.Room_price

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM room_prices WHERE id = ?", id).Scan(&room_prices).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})

}

// GET /Room_types

func ListRoom_prices(c *gin.Context) {

	var room_prices []entity.Room_price

	if err := entity.DB().Raw("SELECT * FROM Room_prices").Scan(&room_prices).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_prices})

}
