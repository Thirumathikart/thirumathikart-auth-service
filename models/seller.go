package models

import (
	"gorm.io/gorm"
)

type Seller struct {
	gorm.Model
	UserID      uint   `gorm:"not null;"`
	Membercode  string `gorm:"default:null;"`
	Ifsccode    string `gorm:"default:null;"`
	Accountno   string `gorm:"default:null;"`
	Accountname string `gorm:"default:null;"`
	Branch      string `gorm:"default:null;"`
}
