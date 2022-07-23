package repository

import (
	"RestaurantService/models"
	"errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
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

func (repo *Repository) FindAll(r *http.Request) ([]models.RestaurantDTO, int64, error) {
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

func (repo *Repository) SaveRestaurant(restaurantDTO *models.RestaurantDTO) (*models.Restaurant, error) {
	var restaurant models.Restaurant = restaurantDTO.ToRestaurant()
	result := repo.db.Create(restaurant)

	if result.Error != nil {
		return nil, errors.New("error while saving restaurant")
	}

	return &restaurant, nil
}

func (repo *Repository) UpdateRestaurant(updatedRestaurant *models.Restaurant) (*models.Restaurant, error) {
	result := repo.db.Save(updatedRestaurant)

	if result.Error != nil {
		return updatedRestaurant, errors.New("error while updating restaurant")
	}

	return updatedRestaurant, nil
}

func (repo *Repository) DeleteRestaurant(id uint) error {
	result := repo.db.Where("id = ?", id).Delete(&models.Restaurant{})

	return result.Error
}
