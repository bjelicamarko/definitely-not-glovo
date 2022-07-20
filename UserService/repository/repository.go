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

func (repo *Repository) CreateUser(newUserDTO *models.NewUserDTO) (*models.User, error) {
	var user models.User

	user.Model = gorm.Model{}
	user.Email = newUserDTO.Email
	user.Password = newUserDTO.Password
	user.FirstName = newUserDTO.FirstName
	user.LastName = newUserDTO.LastName
	user.Contact = newUserDTO.Contact
	user.Banned = false
	user.Role = models.APPUSER

	result := repo.db.Table("users").Create(&user)

	if result.Error != nil {
		return &user, result.Error
	}

	return &user, nil
}
