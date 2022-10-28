package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /Room_types

func CreateRoom_type(c *gin.Context) {

	var room_types entity.Room_type

	if err := c.ShouldBindJSON(&room_types); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&room_types).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_types})

}

// GET /Room_type/:id

func GetRoom_type(c *gin.Context) {

	var room_types entity.Room_type

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Room_types WHERE id = ?", id).Scan(&room_types).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_types})

}

// GET /Room_types

func ListRoom_types(c *gin.Context) {

	var room_types []entity.Room_type

	if err := entity.DB().Raw("SELECT * FROM Room_types").Scan(&room_types).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": room_types})

}
