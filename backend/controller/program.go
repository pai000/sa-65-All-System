package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /job_position
func CreateProgram(c *gin.Context) {
	var program entity.Program
	if err := c.ShouldBindJSON(&program); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&program).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": program})
}

// GET /job_position
// List all job_position
func ListProgram(c *gin.Context) {
	var program []entity.Program
	if err := entity.DB().Raw("SELECT * FROM programs").Scan(&program).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": program})
}

// GET /job_position/:id
// Get job_position by id
func GetProgram(c *gin.Context) {
	var program entity.Program
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&program); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "program not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": program})
}
