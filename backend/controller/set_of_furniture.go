package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /Set_of_furnitures

func CreateSet_of_furniture(c *gin.Context) {

	var set_of_furnitures entity.Set_of_furniture

	if err := c.ShouldBindJSON(&set_of_furnitures); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&set_of_furnitures).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": set_of_furnitures})

}

// GET /Set_of_furniture/:id

func GetSet_of_furniture(c *gin.Context) {

	var set_of_furnitures entity.Set_of_furniture

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Set_of_furnitures WHERE id = ?", id).Scan(&set_of_furnitures).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": set_of_furnitures})

}

// GET /Set_of_furnitures

func ListSet_of_furnitures(c *gin.Context) {

	var set_of_furnitures []entity.Set_of_furniture

	if err := entity.DB().Raw("SELECT * FROM Set_of_furnitures").Scan(&set_of_furnitures).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": set_of_furnitures})

}