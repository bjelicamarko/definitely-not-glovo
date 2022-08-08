package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	IdRestaurant uint `gorm:"not null"`
	IdAppUser    uint `gorm:"not null"`
	IdEmployee   uint
	IdDeliverer  uint
	OrderStatus  OrderStatus `gorm:"not null"`
	TotalPrice   float32     `gorm:"not null;min:0.0"`
	Tip          float32     `gorm:"min:0.0"`
	Note         string      `gorm:"not null"`
	DateTime     string      `gorm:"not null"`
	Country      string      `gorm:"not null"`
	City         string      `gorm:"not null"`
	Street       string      `gorm:"not null"`
	StreetNumber string      `gorm:"not null"`
	Ptt          uint        `gorm:"not null"`
	DisplayName  string      `gorm:"not null"`
	Longitude    float32     `gorm:"not null"`
	Latitude     float32     `gorm:"not null"`
	Reviewed     bool        `gorm:"not null"`
}

type OrderItem struct {
	gorm.Model
	IdOrder      uint    `gorm:"not null"`
	IdArticle    uint    `gorm:"not null"`
	ArticleName  string  `gorm:"not null"`
	CurrentPrice float32 `gorm:"not null;min:0.0"`
	Quantity     uint    `gorm:"not null;minimum(1)"`
	TotalPrice   float32 `gorm:"not null;min:0.0"`
}
