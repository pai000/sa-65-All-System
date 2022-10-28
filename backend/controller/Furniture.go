package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /Room_types
func CreateFurniture(c *gin.Context) {

	var furniture entity.Furniture
	var set_of_furniture entity.Set_of_furniture

	if err := c.ShouldBindJSON(&furniture); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", furniture.Set_of_furniture_id).First(&set_of_furniture); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "set of furniture หาไม่เจอเวนนนน"})

		return

	}

	wv := entity.Furniture{

		Set_of_furniture: set_of_furniture,
	}
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

func GetFurniture(c *gin.Context) {

	var furniture entity.Furniture

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM furnitures WHERE id = ?", id).Scan(&furniture).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": furniture})

}

// GET /Room_types

func ListFurniture(c *gin.Context) {

	var furniture []entity.Furniture

	if err := entity.DB().Raw("SELECT * FROM furnitures").Scan(&furniture).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": furniture})

}
