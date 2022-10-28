package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /semesters
func CreateSemester(c *gin.Context) {
	var semester entity.Semester
	if err := c.ShouldBindJSON(&semester); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&semester).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": semester})
}

// GET /semester/:id
func GetSemester(c *gin.Context) {
	var semester entity.Semester
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&semester); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "semester not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": semester})
}

// GET /semesters
func ListSemesters(c *gin.Context) {
	var semesters []entity.Semester
	if err := entity.DB().Raw("SELECT * FROM semesters").Scan(&semesters).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": semesters})
}
