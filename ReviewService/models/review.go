package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Comment              string `gorm:"not null"`
	Rating               int    `gorm:"not null"`
	InappropriateContent bool   `gorm:"not null"`
	DateTime             string `gorm:"not null"`
	IdRestaurant         uint   `gorm:"not null"`
	IdOrder              uint   `gorm:"not null"`
	IdUser               uint   `gorm:"not null"`
	EmailUser            string `gorm:"not null"`
}
