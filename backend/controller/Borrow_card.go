package controller

import (
	"github.com/pai000/sa-65-project/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Borrow_cards

func CreateBorrow_card(c *gin.Context) {

	var borrow_cards entity.Borrow_card

	if err := c.ShouldBindJSON(&borrow_cards); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&borrow_cards).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": borrow_cards})

}

// GET /Borrow_card/:id

func GetBorrow_card(c *gin.Context) {

	var borrow_cards entity.Borrow_card

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Borrow_cards WHERE id = ?", id).Scan(&borrow_cards).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": borrow_cards})

}

// GET /Borrow_cards

func ListBorrow_cards(c *gin.Context) {

	var borrow_cards []entity.Room_type

	if err := entity.DB().Raw("SELECT * FROM Borrow_cards").Scan(&borrow_cards).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": borrow_cards})

}
