package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pai000/sa-65-project/entity"
)

// POST /Rooms

func CreateRoom(c *gin.Context) {

	var room entity.Room
	var room_price entity.Room_price
	var room_type entity.Room_type
	var set_of_furniture entity.Set_of_furniture

	if err := c.ShouldBindJSON(&room); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", room.Room_type_id).First(&room_type); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room type not found"})
		return

	}
	if tx := entity.DB().Where("id = ?", room.Room_price_id).First(&room_price); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room price not found"})
		return

	}
	if tx := entity.DB().Where("id = ?", room.Set_of_furniture_id).First(&set_of_furniture); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "set of furniture not found"})

		return

	}
	wv := entity.Room{
		Room_type:        room_type,
		Room_price:       room_price,
		Set_of_furniture: set_of_furniture,
	}
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /Room/:id
// GET /watchvideo/:id
func GetRoom(c *gin.Context) {
	var Room entity.Room
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&Room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Room})
}

// GET /watch_videos
func ListRoom(c *gin.Context) {
	var Room []entity.Room
	if err := entity.DB().Preload("Room_type").Preload("Room_price").Preload("Set_of_furniture").Raw("SELECT * FROM Rooms").Find(&Room).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Room})
}
