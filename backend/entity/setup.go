package entity

import (
	//"golang.org/x/crypto/bcrypt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
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
		// &User{},
		&Equipment{},
		&Borrow_card{},
		//&Set_of_furniture{},
		&Room{},
		&List_data{},
		//

		// &Gender{},
		// &Job_Position{},
		// &Province{},
		// &Employee{},
	)

	//password1, err := bcrypt.GenerateFromPassword([]byte("abc12456"), 14)
	//password2, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	//password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)
	password4, err := bcrypt.GenerateFromPassword([]byte("adas8485"), 14)
	//password5, err := bcrypt.GenerateFromPassword([]byte("123"), 14)
	password6, err := bcrypt.GenerateFromPassword([]byte("000"), 14)

	db = database

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

	// --- Room Data
	type1 := Room_type{
		Room_type_name: "Single",
	}
	db.Model(&Room_type{}).Create(&type1)

	type2 := Room_type{
		Room_type_name: "Twin",
	}

	db.Model(&Room_type{}).Create(&type2)

	//Room price Data
	price1 := Room_price{
		Price: "3000",
	}
	db.Model(&Room_price{}).Create(&price1)

	price2 := Room_price{
		Price: "5000",
	}
	db.Model(&Room_price{}).Create(&price2)

	//set_of
	set1 := Set_of_furniture{
		Set_of_furniture_title: "Set1",
	}
	db.Model(&Set_of_furniture{}).Create(&set1)
	set2 := Set_of_furniture{
		Set_of_furniture_title: "Set2",
	}
	db.Model(&Set_of_furniture{}).Create(&set1)

	furniture1 := Furniture{
		Furniture_type: "Table1",

		Set_of_furniture: set1,
	}
	db.Model(&Furniture{}).Create(&furniture1)

	furniture2 := Furniture{
		Furniture_type:   "Table2",
		Set_of_furniture: set2,
	}
	db.Model(&Furniture{}).Create(&furniture2)

	db.Model(&Furniture{}).Create(&Furniture{
		Furniture_type:   "table3",
		Set_of_furniture: set1,
	})

	// === Query
	//
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

	std1 := Student{
		STUDENT_NUMBER: "B62457815",
		STUDENT_NAME:   "Supachai srikawe",
		PERSONAL_ID:    "1786542390457",
		Password:       string(password4),
	}
	db.Model(&Student{}).Create(&std1)

	std2 := Student{
		STUDENT_NUMBER: "B61547843",
		STUDENT_NAME:   "Yanisa wisagesak",
		PERSONAL_ID:    "5698231452357",
		Password:       string(password6),
	}
	db.Model(&Student{}).Create(&std2)

}
