package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ArticleName    string  `gorm:"not null"`
	ArticleType    string  `gorm:"not null"`
	Price          float32 `gorm:"not null"`
	Description    string  `gorm:"not null"`
	RestaurantName string  `gorm:"not null"`
	Image          string  `gorm:"not null"`
}
