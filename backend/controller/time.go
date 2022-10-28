package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)


// POST /Time
func CreateTime(c *gin.Context) {
	var time entity.Time
	if err := c.ShouldBindJSON(&time); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&time).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": time})
}

// GET /Time/:id
func GetTime(c *gin.Context) {
	var time entity.Time
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&time); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "time not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": time})
}

// GET /Time
func ListTimes(c *gin.Context) {
	var times []entity.Time
	if err := entity.DB().Raw("SELECT * FROM times").Scan(&times).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": times})
}