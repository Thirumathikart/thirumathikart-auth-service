package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID   uint `gorm:"not null;"`
	User     User
	Line1    string `gorm:"default:null;"`
	Line2    string `gorm:"default:null;"`
	Landmark string `gorm:"default:null;"`
	District string `gorm:"default:null;"`
	State    string `gorm:"default:null;"`
	Pincode  string `gorm:"default:null;"`
}
