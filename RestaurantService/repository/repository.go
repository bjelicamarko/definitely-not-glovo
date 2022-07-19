package repository

import (
	"RestaurantService/models"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (repo *Repository) FindAll() ([]*models.Restaurant, error) {
	var restaurants []*models.Restaurant

	result := repo.db.Find(&restaurants)

	if result.Error != nil {
		return restaurants, result.Error
	}

	return restaurants, nil
}

func (repo *Repository) SaveRestaurant(restaurant *models.Restaurant) (*models.Restaurant, error) {
	result := repo.db.Create(restaurant)

	if result.Error != nil {
		return restaurant, errors.New("error while saving restaurant")
	}

	return restaurant, nil
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
