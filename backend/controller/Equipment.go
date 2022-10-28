package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /Equipments

func CreateEquipment(c *gin.Context) {

	var equipments entity.Equipment

	if err := c.ShouldBindJSON(&equipments); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&equipments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})

}

// GET /Equipment/:id

func GetEquipment(c *gin.Context) {

	var equipments entity.Equipment

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Equipments WHERE id = ?", id).Scan(&equipments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})

}

// GET /Equipments

func ListEquipments(c *gin.Context) {

	var equipments []entity.Equipment

	if err := entity.DB().Raw("SELECT * FROM Equipments").Scan(&equipments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})

}
