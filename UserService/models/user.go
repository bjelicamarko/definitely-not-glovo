package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"not null;unique_index"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Contact   string `gorm:"not null"`
	Role      Role   `gorm:"not null"`
	Banned    bool   `gorm:"not null"`
	Image     string `gorm:"not null"`
}
