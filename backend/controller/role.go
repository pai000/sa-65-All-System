package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /provinces
func CreateRole(c *gin.Context) {
	var role entity.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": role})
}

// GET /province/:id
func GetRole(c *gin.Context) {
	var province entity.Province
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": province})
}

// GET /provinces
func ListRole(c *gin.Context) {
	var role []entity.Role
	if err := entity.DB().Raw("SELECT * FROM Provinces").Scan(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}
