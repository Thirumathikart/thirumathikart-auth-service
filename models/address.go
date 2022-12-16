package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID   	uint 		`gorm:"not null;"`
	User     	User
	Line1    	string 		`gorm:"not null;"`
	Line2    	string 		`gorm:"default:null;"`
	Landmark 	string 		`gorm:"default:null;"`
	District 	string 		`gorm:"not null;"`
	State    	string 		`gorm:"not null;"`
	Pincode  	string 		`gorm:"not null;"`
	Latitude 	float64 	`gorm:"not null;"`
	Longitude 	float64 	`gorm:"not null;"`
}
