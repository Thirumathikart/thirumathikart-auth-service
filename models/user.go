package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname    string `gorm:"default:null;"`
	Lastname     string `gorm:"default:null;"`
	Email        string `gorm:"not null;unique"`
	Contactno    string `gorm:"default:null;"`
	Passwordhash string `gorm:"not null;"`
	IsVerified   bool   `gorm:"default:false;"`
	Fcm_token    string `gorm:"default:null;"`
	IsSeller     bool   `gorm:"default:false;"`
}
