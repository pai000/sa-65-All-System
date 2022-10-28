package entity

import (
	"gorm.io/gorm"
)

type Room_type struct {
	gorm.Model
	Room_type_name string
	Room           []Room `gorm:"foreignKey:Room_type_id"`
}
type Room_price struct {
	gorm.Model
	Price string

	Room []Room `gorm:"foreignKey:Room_price_id"`
}
type Furniture struct {
	gorm.Model
	Furniture_type string

	Set_of_furniture_id *uint
	Set_of_furniture    Set_of_furniture `gorm:"refernces:id"`
}
type Set_of_furniture struct {
	gorm.Model
	Set_of_furniture_title string

	Furniture []Furniture `gorm:"foreignKey:Set_of_furniture_id"`
	Room      []Room      `gorm:"foreignKey:Set_of_furniture_id"`
}
type Room struct {
	gorm.Model
	Room_type_id *uint
	Room_type    Room_type `gorm:"refernces:id"`

	Room_price_id *uint
	Room_price    Room_price `gorm:"refernces:id"`

	Set_of_furniture_id *uint
	Set_of_furniture    Set_of_furniture `gorm:"refernces:id"`

	Booking   []Booking   `gorm:"foreignKey:Room_id"`
	Repair    []Repair    `gorm:"foreignKey:Room_id"`
	List_data []List_data `gorm:"foreignKey:Room_id"`
}
