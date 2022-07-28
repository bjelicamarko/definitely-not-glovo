package db

import (
	"RestaurantService/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var restaurants = []models.Restaurant{
	{
		Model:          gorm.Model{},
		RestaurantName: "McDonalds",
		ContactPhone:   "05214321",
		Image:          "images/default.png",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Luka Petrovic",
		StreetNumber:   "4",
		DisplayName:    "Trebinje Luka Petrovic 4",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
	},
	{
		Model:          gorm.Model{},
		RestaurantName: "KFC",
		ContactPhone:   "05214321",
		Image:          "images/default.png",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Nikola Tesla",
		StreetNumber:   "3",
		DisplayName:    "Trebinje Nikola Tesla 3",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
	},
	{
		Model:          gorm.Model{},
		RestaurantName: "FOODIC",
		ContactPhone:   "05214321",
		Image:          "images/default.png",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Vojvoda Stepa",
		StreetNumber:   "10",
		DisplayName:    "Trebinje Vojvoda Stepa 10",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
	},
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=restaurantsDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("restaurants")
	db.Migrator().AutoMigrate(&models.Restaurant{})

	for _, restaurant := range restaurants {
		db.Create(&restaurant)
	}

	return db
}
