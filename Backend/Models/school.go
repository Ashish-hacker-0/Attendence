package models

import (
	"gorm.io/gorm"
)

type Attendence struct{
	ID 			 uint		`gorm:"primary key;autoIncrement" json:"id"`
	AID          uint   	`json:"aid"`
	Day      	string    	`json:"day"`
	Month 		string    	`json:"month"`
	Year		string      `json:"year"`
	Punchin		string      `json:"punchin"`
	Punchout		string		`json:"punchout"`
	Type		string		`json:"type"`
	Class		string		`json:"class"`
}

type Student struct{
	SID          	uint  			`gorm:"primary key;autoIncrement" json:"sid"`
	Name    		string  		`json:"name"`
	Class			string			`json:"class"`
}

type Teacher struct{
	TID 				uint		    `gorm:"primary key;autoIncrement" json:"tid"`
	Name    		string  		    `json:"name"`
}

func MigrateStudent(db *gorm.DB) error{
	err:=db.AutoMigrate(&Student{})
	return err
}
func MigrateTeacher(db *gorm.DB) error{
	err:=db.AutoMigrate(&Teacher{})
	return err
}
func MigrateAttendence(db *gorm.DB) error{
	err:=db.AutoMigrate(&Attendence{})
	return err
}
