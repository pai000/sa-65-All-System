package entity

import (
	"time"

	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model

	Semester     string
	Payment_Bill []Payment_Bill `gorm:"foreignKey:SemesterID"`
}

type Payment_Bill struct {
	gorm.Model
	Billing_Date    time.Time
	Electric_Bill   float32
	Water_Bill      float32
	Payment_Balance float32

	// Employee ทำหน้าที่เป็น FK
	EmployeeID int
	Employee   Employee `gorm:"references:id"`

	// BookingID ทำหน้าที่เป็น FK
	BookingID int
	Booking   Booking `gorm:"references:id"`

	// Semester ทำหน้าที่เป็น FK
	SemesterID int
	Semester   Semester `gorm:"references:id"`
}
