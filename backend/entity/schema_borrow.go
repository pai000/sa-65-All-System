package entity

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	gorm.Model
	Equipment_name        string
	Equipment_code        string
	Equipment_borrower_id *uint       `gorm:"foreignKey:Student_id"`
	List_data             []List_data `gorm:"foreignKey:Equipment_id"`
}

type Borrow_card struct {
	gorm.Model
	//Price string
	Borrower_id *uint

	List_data []List_data `gorm:"foreignKey:Borrow_card_id"`
}

type List_data struct {
	gorm.Model
	Borrow_time  time.Time
	Return_time  time.Time
	Equipment_id *uint
	Equipment    Equipment `gorm:"refernces:id"`

	Borrow_card_id *uint
	Borrow_card    Borrow_card `gorm:"refernces:id"`

	Room_id *uint
	Room    Room `gorm:"refernces:id"`
}
