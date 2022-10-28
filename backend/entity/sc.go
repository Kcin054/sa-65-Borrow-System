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

	List_data []List_data `gorm:"foreignKey:Room_id"`
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

// /
type Gender struct {
	gorm.Model
	Name string

	Employees []Employee `gorm:"foreignKey:GenderID"`
}

type Job_Position struct {
	gorm.Model
	Name string

	Employees []Employee `gorm:"foreignKey:Job_PositionID"`
}

type Province struct {
	gorm.Model
	Name string

	Employees []Employee `gorm:"foreignKey:ProvinceID"`
}

type Employee struct {
	gorm.Model
	Personal_ID string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Name        string
	Password    string

	//GenderID ทำหน้าที่เป็น FK
	GenderID *uint
	Gender   Gender `gorm:"references:id"`

	//Job_PositionID ทำหน้าที่เป็น FK
	Job_PositionID *uint
	Job_Position   Job_Position `gorm:"references:id"`

	//ProvinceID ทำหน้าที่เป็น FK
	ProvinceID *uint
	Province   Province `gorm:"references:id"`
}
