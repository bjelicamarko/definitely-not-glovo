package repository

import (
	"UserService/models"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func concat(str string) string {
	return "%" + strings.ToLower(str) + "%"
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 0 {
			page = 0
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
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

func (repo *Repository) FindAll(r *http.Request) ([]models.UserDTO, int64, error) {
	var usersDTO []models.UserDTO

	var users []*models.User

	var totalElements int64

	result := repo.db.Scopes(Paginate(r)).Table("users").Where("role != ?", models.ADMIN).Find(&users)
	repo.db.Table("users").Where("role != ? and deleted_at IS NULL", models.ADMIN).Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, user := range users {
		usersDTO = append(usersDTO, user.ToUserDTO())
	}

	return usersDTO, totalElements, nil
}

func (repo *Repository) SearchUsers(r *http.Request) ([]models.UserDTO, int64, error) {
	var usersDTO []models.UserDTO
	var users []*models.User
	var totalElements int64

	searchField := r.URL.Query().Get("searchField")
	userType := r.URL.Query().Get("userType")

	result := repo.db.Scopes(Paginate(r)).Table("users").
		Where("(deleted_at IS NULL and role != ?) and "+
			"('' = ? or "+
			"first_name LIKE ? or "+
			"last_name LIKE ? or "+
			"email LIKE ?) and"+
			"('' = ? or role = ?)",
			models.ADMIN, searchField, concat(searchField), concat(searchField), concat(searchField), userType, userType).
		Find(&users)

	repo.db.Table("users").
		Where("(deleted_at IS NULL and role != ?) and "+
			"('' = ? or "+
			"first_name LIKE ? or "+
			"last_name LIKE ? or "+
			"email LIKE ?) and"+
			"('' = ? or role = ?)",
			models.ADMIN, searchField, concat(searchField), concat(searchField), concat(searchField), userType, userType).
		Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, user := range users {
		usersDTO = append(usersDTO, user.ToUserDTO())
	}

	return usersDTO, totalElements, nil
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

func (repo *Repository) UpdateUser(userDTO *models.UserDTO, indicator bool) (*models.User, error) {
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

	if indicator {
		user.Image = userDTO.Image
	}

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

func (repo *Repository) BanUser(id uint) error {
	var user models.User
	result := repo.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return result.Error
	}

	user.Banned = true
	repo.db.Table("users").Save(&user)

	return nil
}

func (repo *Repository) UnbanUser(id uint) error {
	var user models.User
	result := repo.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return result.Error
	}

	user.Banned = false
	repo.db.Table("users").Save(&user)

	return nil
}

func (repo *Repository) FindUserById(id uint) (*models.UserDTO, error) {
	var user models.User

	result := repo.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	var userDTO models.UserDTO = user.ToUserDTO()
	return &userDTO, nil
}
