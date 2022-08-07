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
		Comment:              "Odlicno",
		Rating:               10,
		InappropriateContent: false,
		DateTime:             "08.07.2022. 12:16",
		IdRestaurant:         1,
		IdOrder:              1,
		IdUser:               2,
		EmailUser:            "pera@gmail.com",
	},
	{
		Model:                gorm.Model{},
		Comment:              "Super",
		Rating:               10,
		InappropriateContent: false,
		DateTime:             "08.07.2022. 12:16",
		IdRestaurant:         1,
		IdOrder:              2,
		IdUser:               2,
		EmailUser:            "pera@gmail.com",
	},
	{
		Model:                gorm.Model{},
		Comment:              "Dobro je",
		Rating:               10,
		InappropriateContent: false,
		DateTime:             "08.07.2022. 12:16",
		IdRestaurant:         1,
		IdOrder:              3,
		IdUser:               2,
		EmailUser:            "pera@gmail.com",
	},
	{
		Model:                gorm.Model{},
		Comment:              "Drugi restoran  bolji",
		Rating:               5,
		InappropriateContent: false,
		DateTime:             "08.07.2022. 12:16",
		IdRestaurant:         2,
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
