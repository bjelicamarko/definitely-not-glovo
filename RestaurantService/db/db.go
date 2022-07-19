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
		Street:         "Milorad Lecic",
		StreetNumber:   "4",
		City:           "Trebinje",
		ContactPhone:   "05214321",
	},
	{
		Model:          gorm.Model{},
		RestaurantName: "KFC",
		Street:         "Ulica Nikole Tesle",
		StreetNumber:   "3",
		City:           "Trebinje",
		ContactPhone:   "05214321",
	},
	{
		Model:          gorm.Model{},
		RestaurantName: "Foodic",
		Street:         "Bulevar Despota Stefana",
		StreetNumber:   "7a",
		City:           "Novi Sad",
		ContactPhone:   "05214321",
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
