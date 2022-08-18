package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	RestaurantName string  `gorm:"not null;unique"`
	ContactPhone   string  `gorm:"not null"`
	Image          string  `gorm:"not null"`
	Country        string  `gorm:"not null"`
	City           string  `gorm:"not null"`
	Street         string  `gorm:"not null"`
	StreetNumber   string  `gorm:"not null"`
	Ptt            uint    `gorm:"not null"`
	DisplayName    string  `gorm:"not null"`
	Longitude      float32 `gorm:"not null"`
	Latitude       float32 `gorm:"not null"`
	Delivery       float32 `gorm:"not null"`
}
