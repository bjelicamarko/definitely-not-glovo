package repository

import (
	"UserService/models"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo *Repository) FindAllUsers(r *http.Request) ([]models.UserDTO, int64, error) {
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
			"lower(first_name) LIKE ? or "+
			"lower(last_name) LIKE ? or "+
			"lower(email) LIKE ?) and"+
			"('' = ? or role = ?)",
			models.ADMIN, searchField, concat(searchField), concat(searchField), concat(searchField), userType, userType).
		Find(&users)

	repo.db.Table("users").
		Where("(deleted_at IS NULL and role != ?) and "+
			"('' = ? or "+
			"lower(first_name) LIKE ? or "+
			"lower(last_name) LIKE ? or "+
			"lower(email) LIKE ?) and"+
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

func (repo *Repository) FindUserById(id uint) (*models.UserDTO, error) {
	var user models.User
	result := repo.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("user cannot be found")
	}

	var userDTO models.UserDTO = user.ToUserDTO()
	return &userDTO, nil
}

func (repo *Repository) Register(userDTO *models.UserDTO) error {
	var user models.User

	user.Model = gorm.Model{}
	user.Email = userDTO.Email
	user.Password = userDTO.Password
	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.Contact = userDTO.Contact
	user.Role = models.APPUSER
	user.Banned = false
	user.Image = "images/default.jpg"

	result := repo.db.Table("users").Create(&user)
	return result.Error
}

func (repo *Repository) CreateUser(userDTO *models.UserDTO) (*models.UserDTO, error) {
	var user models.User = userDTO.ToUser()
	result := repo.db.Table("users").Create(&user)

	if result.Error != nil {
		return nil, errors.New("error while creating user")
	}

	var retValue models.UserDTO = user.ToUserDTO()
	return &retValue, nil
}

func (repo *Repository) UpdateUser(userDTO *models.UserDTO) (*models.UserDTO, error) {
	var user models.User
	result := repo.db.Table("users").Where("id = ?", userDTO.Id).First(&user)

	if result.Error != nil {
		return nil, errors.New("user cannot be found")
	}

	if userDTO.Changed {
		user.Image = userDTO.ImagePath
	}

	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.Contact = userDTO.Contact
	user.Banned = userDTO.Banned

	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("error while updating user")
	}

	var retValue models.UserDTO = user.ToUserDTO()
	return &retValue, nil
}

func (repo *Repository) DeleteUser(id uint) (*models.UserDTO, error) {
	var user models.User
	result := repo.db.Where("id = ?", id).Clauses(clause.Returning{}).Delete(&user)

	if result.Error != nil {
		return nil, errors.New("error while deleting user")
	}

	var userDTO models.UserDTO = user.ToUserDTO()
	return &userDTO, nil
}

func (repo *Repository) BanUser(id uint) (*models.UserDTO, error) {
	var user models.User
	result := repo.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("user cannot be found")
	}

	user.Banned = true
	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("error while banning user")
	}

	var userDTO models.UserDTO = user.ToUserDTO()
	return &userDTO, nil
}

func (repo *Repository) UnbanUser(id uint) (*models.UserDTO, error) {
	var user models.User
	result := repo.db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		return nil, errors.New("user cannot be found")
	}

	user.Banned = false
	result2 := repo.db.Table("users").Save(&user)

	if result2.Error != nil {
		return nil, errors.New("error while unbanning user")
	}

	var userDTO models.UserDTO = user.ToUserDTO()
	return &userDTO, nil
}
