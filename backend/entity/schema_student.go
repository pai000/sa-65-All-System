package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Role_name string
	Students  []Student `gorm:"foreignKey:RoleID"`
}
type Program struct {
	gorm.Model
	Program_name string
	Students     []Student `gorm:"foreignKey:ProgramID"`
}
type Student struct {
	gorm.Model
	STUDENT_NUMBER string `gorm:"uniqueIndex"`
	STUDENT_NAME   string
	PERSONAL_ID    string `gorm:"uniqueIndex"`
	Password       string

	//GenderID ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id"`
	//ProvinceID ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id"`
	//ProgramID ทำหน้าที่เป็น FK
	ProgramID *uint
	Program   Program `gorm:"references:id"`
	//RoleID ทำหน้าที่เป็น FK
	RoleID *uint
	Role   Role `gorm:"references:id"`
	//EmpolyeeID ทำหน้าที่เป็น FK
	EmployeeID *uint
	Employee   Employee
	Booking    []Booking `gorm:"foreignKey:STUDENT_ID"`
	Repair     []Repair  `gorm:"foreignKey:STUDENT_ID"`
}
