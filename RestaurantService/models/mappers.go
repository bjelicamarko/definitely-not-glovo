package models

import (
	"RestaurantService/utils"

	"gorm.io/gorm"
)

func (restaurant *Restaurant) ToRestaurantDTO() RestaurantDTO {
	return RestaurantDTO{
		Id:             restaurant.ID,
		RestaurantName: restaurant.RestaurantName,
		ContactPhone:   restaurant.ContactPhone,
		Image:          utils.GetB64Image(restaurant.Image),
		ImagePath:      restaurant.Image,
		Country:        restaurant.Country,
		City:           restaurant.City,
		Street:         restaurant.Street,
		StreetNumber:   restaurant.StreetNumber,
		Ptt:            restaurant.Ptt,
		DisplayName:    restaurant.DisplayName,
		Longitude:      restaurant.Longitude,
		Latitude:       restaurant.Latitude,
		Changed:        false,
	}
}

func (restaurantDTO *RestaurantDTO) ToRestaurant() Restaurant {
	return Restaurant{
		Model:          gorm.Model{},
		RestaurantName: restaurantDTO.RestaurantName,
		ContactPhone:   restaurantDTO.ContactPhone,
		Image:          restaurantDTO.ImagePath,
		Country:        restaurantDTO.Country,
		City:           restaurantDTO.City,
		Street:         restaurantDTO.Street,
		StreetNumber:   restaurantDTO.StreetNumber,
		Ptt:            restaurantDTO.Ptt,
		DisplayName:    restaurantDTO.DisplayName,
		Longitude:      restaurantDTO.Longitude,
		Latitude:       restaurantDTO.Latitude,
	}
}
