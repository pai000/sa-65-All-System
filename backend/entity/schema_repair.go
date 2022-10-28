package entity

import (
	"gorm.io/gorm"
)

// /ระบบแจ้งซ่อม
// Repair
type Repair struct {
	gorm.Model
	Comment      string
	STUDENT_ID   *uint
	Student      Student `gorm:"refernces:STUDENT_ID"`
	Room_id      *uint
	Room         Room `gorm:"refernces:Room_id"`
	Furniture_id *uint
	Furniture    Furniture `gorm:"refernces:furniture_id"`
}
