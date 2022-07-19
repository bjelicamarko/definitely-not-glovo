package repository

import (
	"UserService/models"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) CheckCredentials(email string, password string) (*models.User, error) {
	var user models.User

	repo.db.Table("users").Where("email = ? AND password = ?", email, password).Find(&user)

	if user.ID == 0 {
		return &user, errors.New("invalid username or password")
	}

	if user.Banned {
		return &user, errors.New("user is banned")
	}

	return &user, nil
}
