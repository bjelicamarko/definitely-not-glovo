package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	RestaurantName string `gorm:"not null;unique"`
	Street         string `gorm:"not null"`
	StreetNumber   string `gorm:"not null"`
	City           string `gorm:"not null"`
	ContactPhone   string `gorm:"not null"`
}
