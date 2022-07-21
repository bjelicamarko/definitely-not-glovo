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

func (repo *Repository) FindAll() ([]models.UserDTO, error) {
	var usersDTO []models.UserDTO

	var users []*models.User

	result := repo.db.Table("users").Where("role != ?", models.ADMIN).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	for _, user := range users {
		usersDTO = append(usersDTO, user.ToUserDTO())
	}

	return usersDTO, nil
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

func (repo *Repository) UpdateUser(userDTO *models.UserDTO) (*models.User, error) {
	var user models.User

	result := repo.db.Table("users").Where("email = ?", userDTO.Email).First(&user)

	if result.Error != nil {
		return nil, errors.New("user cannot be found")
	}

	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.Contact = userDTO.Contact
	user.Banned = userDTO.Banned
	user.Role = models.Role(userDTO.Role)

	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("error while updating user")
	}

	return &user, nil
}

func (repo *Repository) DeleteUser(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&models.User{})

	return result.Error
}
