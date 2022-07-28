package repository

import (
	"RestaurantService/models"
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

func (repo *Repository) FindAllRestaurants(r *http.Request) ([]models.RestaurantDTO, int64, error) {
	var restaurantsDTO []models.RestaurantDTO
	var restaurants []*models.Restaurant
	var totalElements int64

	result := repo.db.Scopes(Paginate(r)).Table("restaurants").Find(&restaurants)
	repo.db.Table("restaurants").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, restaurant := range restaurants {
		restaurantsDTO = append(restaurantsDTO, restaurant.ToRestaurantDTO())
	}

	return restaurantsDTO, totalElements, nil
}

func (repo *Repository) SearchRestaurants(r *http.Request) ([]models.RestaurantDTO, int64, error) {
	var restaurantsDTO []models.RestaurantDTO
	var restaurants []*models.Restaurant
	var totalElements int64

	searchField := r.URL.Query().Get("searchField")

	result := repo.db.Scopes(Paginate(r)).Table("restaurants").
		Where("(deleted_at IS NULL) and "+
			"('' = ? or lower(restaurant_name) LIKE ?)",
			searchField, concat(searchField)).
		Find(&restaurants)

	repo.db.Table("restaurants").
		Where("(deleted_at IS NULL) and "+
			"('' = ? or lower(restaurant_name) LIKE ?)",
			searchField, concat(searchField)).
		Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, restaurant := range restaurants {
		restaurantsDTO = append(restaurantsDTO, restaurant.ToRestaurantDTO())
	}

	return restaurantsDTO, totalElements, nil
}

func (repo *Repository) FindRestaurantById(id uint) (*models.RestaurantDTO, error) {
	var restaurant models.Restaurant
	result := repo.db.Table("restaurants").Where("id = ?", id).First(&restaurant)

	if result.Error != nil {
		return nil, errors.New("restaurant cannot be found")
	}

	var restaurantDTO models.RestaurantDTO = restaurant.ToRestaurantDTO()
	return &restaurantDTO, nil
}

func (repo *Repository) CreateRestaurant(restaurantDTO *models.RestaurantDTO) (*models.RestaurantDTO, error) {
	var restaurant models.Restaurant = restaurantDTO.ToRestaurant()
	result := repo.db.Table("restaurants").Create(&restaurant)

	if result.Error != nil {
		return nil, errors.New("error while creating restaurant")
	}

	var retValue models.RestaurantDTO = restaurant.ToRestaurantDTO()
	return &retValue, nil
}

func (repo *Repository) UpdateRestaurant(restaurantDTO *models.RestaurantDTO) (*models.RestaurantDTO, error) {
	var restaurant models.Restaurant
	result := repo.db.Table("restaurants").Where("id = ?", restaurantDTO.Id).First(&restaurant)

	if result.Error != nil {
		return nil, errors.New("restaurant cannot be found")
	}

	if restaurantDTO.Changed {
		restaurant.Image = restaurantDTO.ImagePath
	}

	restaurant.ContactPhone = restaurantDTO.ContactPhone
	restaurant.Country = restaurantDTO.Country
	restaurant.City = restaurantDTO.City
	restaurant.Street = restaurantDTO.Street
	restaurant.StreetNumber = restaurantDTO.StreetNumber
	restaurant.Ptt = restaurantDTO.Ptt
	restaurant.DisplayName = restaurantDTO.DisplayName
	restaurant.Longitude = restaurantDTO.Longitude
	restaurant.Latitude = restaurantDTO.Latitude

	result2 := repo.db.Table("restaurants").Save(&restaurant)

	if result2.Error != nil {
		return nil, errors.New("error while updating restaurant")
	}

	var retValue models.RestaurantDTO = restaurant.ToRestaurantDTO()
	return &retValue, nil
}

func (repo *Repository) DeleteRestaurant(id uint) (*models.RestaurantDTO, error) {
	var restaurant models.Restaurant
	result := repo.db.Where("id = ?", id).Clauses(clause.Returning{}).Delete(&restaurant)

	if result.Error != nil {
		return nil, errors.New("error while deleting restaurant")
	}

	var retValue models.RestaurantDTO = restaurant.ToRestaurantDTO()
	return &retValue, nil
}
