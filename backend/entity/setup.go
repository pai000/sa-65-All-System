package entity

import (
	//"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		//================ Employee =====================
		//===============================================
		&Gender{},
		&Job_Position{},
		&Province{},
		&Employee{},

		//================ Student =====================
		//==============================================
		&Role{},
		&Program{},
		&Student{},

		//================== Room ======================
		//==============================================
		&Room_type{},
		&Room_price{},
		&Set_of_furniture{},
		&Furniture{},
		&Room{},

		//================== Repair ====================
		//==============================================
		&Repair{},

		//================ Booking =====================
		//==============================================
		&Time{},
		&Booking{},

		//================= Borrow =====================
		//==============================================
		&Equipment{},
		&Borrow_card{},
		&List_data{},
	)

	db = database

	//add example

	password1, err := bcrypt.GenerateFromPassword([]byte("abc12456"), 14)
	password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)
	password4, err := bcrypt.GenerateFromPassword([]byte("adas8485"), 14)
	password5, err := bcrypt.GenerateFromPassword([]byte("123"), 14)
	password6, err := bcrypt.GenerateFromPassword([]byte("000"), 14)

	// ======================================================================================================================
	// ======================================  Employee  ====================================================================
	// ======================================================================================================================

	//Gender
	gender1 := Gender{
		GENDER_NAME: "Male",
	}

	db.Model(&Gender{}).Create(&gender1)

	gender2 := Gender{
		GENDER_NAME: "FeMale",
	}

	db.Model(&Gender{}).Create(&gender2)

	//insert job_position
	job_position1 := Job_Position{
		Name: "Admin",
	}
	db.Model(&Job_Position{}).Create(&job_position1)

	job_position2 := Job_Position{
		Name: "Housekeeper",
	}
	db.Model(&Job_Position{}).Create(&job_position2)

	job_position3 := Job_Position{
		Name: "Security Guard",
	}
	db.Model(&Job_Position{}).Create(&job_position3)

	job_position4 := Job_Position{
		Name: "Mechanic",
	}
	db.Model(&Job_Position{}).Create(&job_position4)

	//province
	roiet := Province{
		Name: "Roiet",
	}
	db.Model(&Province{}).Create(&roiet)
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//insert employee'
	em1 := Employee{
		Personal_ID: "1456287463254",
		Email:       "ana@gmail.com",
		Name:        "Ana poul",
		Password:    string(password1),

		Gender:       gender2,
		Job_Position: job_position1,
		Province:     korat,
	}
	db.Model(&Employee{}).Create(&em1)

	em2 := Employee{
		Personal_ID: "5874621453054",
		Email:       "kerkkiat@gmail.com",
		Name:        "Kerkkiat Prabmontree",
		Password:    string(password3),

		Gender:       gender1,
		Job_Position: job_position3,
		Province:     bangkok,
	}
	db.Model(&Employee{}).Create(&em2)

	em3 := Employee{
		Personal_ID: "4587652145385",
		Email:       "matinez@gmail.com",
		Name:        "Devid Matinez",
		Password:    string(password2),

		Gender:       gender1,
		Job_Position: job_position4,
		Province:     chon,
	}
	db.Model(&Employee{}).Create(&em3)

	em4 := Employee{
		Personal_ID: "5847532016420",
		Email:       "akira@gmail.com",
		Name:        "Akira komisu",
		Password:    string(password5),

		Gender:       gender1,
		Job_Position: job_position1,
		Province:     roiet,
	}
	db.Model(&Employee{}).Create(&em4)

	// ======================================================================================================================
	// ======================================  Student  =====================================================================
	// ======================================================================================================================

	// --- Program Data
	p1 := Program{
		Program_name: "Computer engineering",
	}
	db.Model(&Program{}).Create(&p1)
	p2 := Program{
		Program_name: "Telecommunication engineering",
	}
	db.Model(&Program{}).Create(&p2)
	p3 := Program{
		Program_name: "Program in Biology",
	}
	db.Model(&Program{}).Create(&p3)
	p4 := Program{
		Program_name: "Institute of Nursing",
	}
	db.Model(&Program{}).Create(&p4)

	// --- Role Data

	role1 := Role{
		Role_name: "Student",
	}
	db.Model(&Role{}).Create(&role1)

	std1 := Student{
		STUDENT_NUMBER: "B62457815",
		STUDENT_NAME:   "Supachai srikawe",
		PERSONAL_ID:    "1786542390457",
		Password:       string(password4),

		Gender:   gender1,
		Program:  p3,
		Province: roiet,
		Role:     role1,
		Employee: em1,
	}
	db.Model(&Student{}).Create(&std1)

	std2 := Student{
		STUDENT_NUMBER: "B61547843",
		STUDENT_NAME:   "Yanisa wisagesak",
		PERSONAL_ID:    "5698231452357",
		Password:       string(password6),

		Gender:   gender2,
		Program:  p3,
		Province: roiet,
		Role:     role1,
		Employee: em1,
	}
	db.Model(&Student{}).Create(&std2)

	// ======================================================================================================================
	// ============================================  Room  ==================================================================
	// ======================================================================================================================

	type1 := Room_type{
		Room_type_name: "Single",
	}
	db.Model(&Room_type{}).Create(&type1)

	type2 := Room_type{
		Room_type_name: "Twin",
	}
	db.Model(&Room_type{}).Create(&type2)
	type3 := Room_type{
		Room_type_name: "Quad",
	}

	db.Model(&Room_type{}).Create(&type3)

	//price Data
	price1 := Room_price{
		Price: "3000",
	}
	db.Model(&Room_price{}).Create(&price1)

	price2 := Room_price{
		Price: "5000",
	}
	db.Model(&Room_price{}).Create(&price2)

	price3 := Room_price{
		Price: "8000",
	}
	db.Model(&Room_price{}).Create(&price3)

	//set_of
	set1 := Set_of_furniture{
		Set_of_furniture_title: "Set1",
	}
	db.Model(&Set_of_furniture{}).Create(&set1)
	set2 := Set_of_furniture{
		Set_of_furniture_title: "Set2",
	}
	db.Model(&Set_of_furniture{}).Create(&set2)
	set3 := Set_of_furniture{
		Set_of_furniture_title: "Set3",
	}

	db.Model(&Set_of_furniture{}).Create(&set3)
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "table1",
		Set_of_furniture: set1,
	})

	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "fan1",
		Set_of_furniture: set1,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "bed1",
		Set_of_furniture: set1,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "chair1",
		Set_of_furniture: set1,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "led lightning1",
		Set_of_furniture: set1,
	})

	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "table2",
		Set_of_furniture: set2,
	})

	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "Air2",
		Set_of_furniture: set2,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "Tv2",
		Set_of_furniture: set2,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "bed2",
		Set_of_furniture: set2,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "chair2",
		Set_of_furniture: set2,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "led lightning2",
		Set_of_furniture: set2,
	})

	// furniture set3
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "table3",
		Set_of_furniture: set3,
	})

	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "Air3",
		Set_of_furniture: set3,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "Tv3",
		Set_of_furniture: set3,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "bed3",
		Set_of_furniture: set3,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "chair3",
		Set_of_furniture: set3,
	})
	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "led lightning3",
		Set_of_furniture: set3,
	})

	room1 := Room{
		Room_type:        type1,
		Room_price:       price1,
		Set_of_furniture: set1,
	}
	db.Model(&Room{}).Create(&room1)

	room2 := Room{
		Room_type:        type2,
		Room_price:       price1,
		Set_of_furniture: set1,
	}
	db.Model(&Room{}).Create(&room2)

	// ======================================================================================================================
	// ============================================ Booking =================================================================
	// ======================================================================================================================

	T1 := Time{
		Time_number: "1 เทอม",
	}
	db.Model(&Time{}).Create(&T1)
	T2 := Time{
		Time_number: "2 เทอม",
	}
	db.Model(&Time{}).Create(&T2)
	T3 := Time{
		Time_number: "1 ปี",
	}
	db.Model(&Time{}).Create(&T3)
	T4 := Time{
		Time_number: "1 ปี 1 เทอม",
	}
	db.Model(&Time{}).Create(&T4)

	r1 := Room{
		Room_type:        type1,
		Room_price:       price1,
		Set_of_furniture: set1,
	}
	db.Model(&Room{}).Create(&r1)

	db.Model(&Room{}).Create(&Room{
		Room_type:        type1,
		Room_price:       price1,
		Set_of_furniture: set1,
	})

	db.Model(&Room{}).Create(&Room{
		Room_type:        type2,
		Room_price:       price1,
		Set_of_furniture: set1,
	})
	db.Model(&Booking{}).Create(&Booking{
		Room: r1,
		Time: T1,
	})

	// ======================================================================================================================
	// =========================================== Borrow ===================================================================
	// ======================================================================================================================

	equipment1 := Equipment{
		Equipment_name: "Stepladder",
		Equipment_code: "12345",
	}
	db.Model(&Equipment{}).Create(&equipment1)

	equipment2 := Equipment{
		Equipment_name: "Screwdriver",
		Equipment_code: "54218",
	}
	db.Model(&Equipment{}).Create(&equipment2)

	equipment3 := Equipment{
		Equipment_name: "Drill",
		Equipment_code: "33254",
	}
	db.Model(&Equipment{}).Create(&equipment3)

	//Borrow Card Data
	borrow_card1 := Borrow_card{}
	db.Model(&Borrow_card{}).Create(&borrow_card1)

	borrow_card2 := Borrow_card{}
	db.Model(&Borrow_card{}).Create(&borrow_card2)

	db.Model(&List_data{}).Create(&List_data{
		Model:          gorm.Model{},
		Borrow_time:    time.Now(),
		Return_time:    time.Now(),
		Equipment_id:   new(uint),
		Equipment:      equipment1,
		Borrow_card_id: new(uint),
		Borrow_card:    borrow_card1,
		Room_id:        new(uint),
		Room:           Room{},
	})

}
