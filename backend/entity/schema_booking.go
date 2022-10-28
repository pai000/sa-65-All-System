package entity

import (
	"time"

	"gorm.io/gorm"
)

// ระบบจองห้อง
type Time struct {
	gorm.Model
	Time_number string
	Booking     []Booking `gorm:"foreignKey:TimeID"`
}

type Booking struct {
	gorm.Model

	Check_in_date time.Time
	Room_id       *uint
	Room          Room    `gorm:"references:id"`
	STUDENT_ID    *uint   `gorm:"uniqueIndex"`
	Student       Student `gorm:"references:id"`
	TimeID        *uint
	Time          Time `gorm:"references:id"`

	//Payment_Bill []Payment_Bill `gorm:"foreignKey:BookingID"`
}
