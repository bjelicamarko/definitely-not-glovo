package db

import (
	"UserService/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Model:          gorm.Model{}, // 1
		Email:          "admin@gmail.com",
		Password:       "admin",
		FirstName:      "Adminko",
		LastName:       "Adminic",
		Contact:        "231321",
		Role:           models.ADMIN,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "",
	},
	{
		Model:          gorm.Model{}, // 2
		Email:          "pera@gmail.com",
		Password:       "admin",
		FirstName:      "Pera",
		LastName:       "Peric",
		Contact:        "231321",
		Role:           models.APPUSER,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "",
	},
	{
		Model:          gorm.Model{}, // 3
		Email:          "zika@gmail.com",
		Password:       "admin",
		FirstName:      "Zika",
		LastName:       "Zikic",
		Contact:        "231321",
		Role:           models.DELIVERER,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "McDonalds",
	},
	{
		Model:          gorm.Model{}, // 4
		Email:          "dunja@gmail.com",
		Password:       "admin",
		FirstName:      "Dunja",
		LastName:       "Dunjica",
		Contact:        "231321",
		Role:           models.EMPLOYEE,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "McDonalds",
	},
	{
		Model:          gorm.Model{}, // 5
		Email:          "visnja@gmail.com",
		Password:       "admin",
		FirstName:      "Visnja",
		LastName:       "Visnjica",
		Contact:        "231321",
		Role:           models.EMPLOYEE,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "McDonalds",
	},
	{
		Model:          gorm.Model{}, // 6
		Email:          "mrvica@gmail.com",
		Password:       "admin",
		FirstName:      "Mrva",
		LastName:       "Mrvica",
		Contact:        "231321",
		Role:           models.EMPLOYEE,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "KFC",
	},
	{
		Model:          gorm.Model{}, // 7
		Email:          "dzoni@gmail.com",
		Password:       "admin",
		FirstName:      "Dzoni",
		LastName:       "Dzonic",
		Contact:        "231321",
		Role:           models.DELIVERER,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "KFC",
	},
	{
		Model:          gorm.Model{}, // 8
		Email:          "calvin@gmail.com",
		Password:       "admin",
		FirstName:      "Calvin",
		LastName:       "Harris",
		Contact:        "231321",
		Role:           models.APPUSER,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "",
	},
	{
		Model:          gorm.Model{}, // 9
		Email:          "dule@gmail.com",
		Password:       "admin",
		FirstName:      "Dule",
		LastName:       "Dule",
		Contact:        "231321",
		Role:           models.EMPLOYEE,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "FOODIC",
	},
	{
		Model:          gorm.Model{}, // 10
		Email:          "branka@gmail.com",
		Password:       "admin",
		FirstName:      "Branka",
		LastName:       "Brankic",
		Contact:        "231321",
		Role:           models.DELIVERER,
		Banned:         false,
		Image:          "images/default.jpg",
		RestaurantName: "FOODIC",
	},
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=appusersDB port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to db")
	} else {
		fmt.Println("Database connection successfully created")
	}

	db.Migrator().DropTable("users")
	db.Migrator().AutoMigrate(&models.User{})

	for _, user := range users {
		db.Create(&user)
	}

	return db
}
