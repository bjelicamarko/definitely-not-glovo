package db

import (
	"ReviewService/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var reviews = []models.Review{
	{
		Model:                gorm.Model{},
		Comment:              "Komentar 1",
		Rating:               10,
		InappropriateContent: false,
		DateTime:             "17.07.2022. 19:00",
		IdRestaurant:         1,
		IdOrder:              1,
		IdUser:               2,
		EmailUser:            "pera@gmail.com",
	},
	{
		Model:                gorm.Model{},
		Comment:              "Komentar 2",
		Rating:               9,
		InappropriateContent: false,
		DateTime:             "18.07.2022. 18:45",
		IdRestaurant:         2,
		IdOrder:              2,
		IdUser:               2,
		EmailUser:            "pera@gmail.com",
	},
	{
		Model:                gorm.Model{},
		Comment:              "Komentar 3",
		Rating:               8,
		InappropriateContent: false,
		DateTime:             "19.07.2022. 19:23",
		IdRestaurant:         3,
		IdOrder:              3,
		IdUser:               2,
		EmailUser:            "pera@gmail.com",
	},
	{
		Model:                gorm.Model{},
		Comment:              "Komentar 4",
		Rating:               8,
		InappropriateContent: false,
		DateTime:             "20.07.2022. 21:00",
		IdRestaurant:         1,
		IdOrder:              4,
		IdUser:               2,
		EmailUser:            "pera@gmail.com",
	},
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=reviewsDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("reviews")

	db.Migrator().AutoMigrate(&models.Review{})

	for _, review := range reviews {
		db.Create(&review)
	}

	return db
}
