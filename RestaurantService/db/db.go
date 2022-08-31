package db

import (
	"RestaurantService/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var restaurants = []models.Restaurant{
	{
		Model:          gorm.Model{},
		RestaurantName: "McDonalds",
		ContactPhone:   "05214321",
		Image:          "images/mcdonalds.png",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Luka Petrovic",
		StreetNumber:   "4",
		DisplayName:    "Trebinje Luka Petrovic 4",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
		Delivery:       200.0,
	},
	{
		Model:          gorm.Model{},
		RestaurantName: "KFC",
		ContactPhone:   "05214321",
		Image:          "images/kfc.png",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Nikola Tesla",
		StreetNumber:   "3",
		DisplayName:    "Trebinje Nikola Tesla 3",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
		Delivery:       250.0,
	},
	{
		Model:          gorm.Model{},
		RestaurantName: "FOODIC",
		ContactPhone:   "05214321",
		Image:          "images/foodic.jpg",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Vojvoda Stepa",
		StreetNumber:   "10",
		DisplayName:    "Trebinje Vojvoda Stepa 10",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
		Delivery:       180.0,
	},
}

func Init() *gorm.DB {
	host := os.Getenv("pgHost")
	port := os.Getenv("pgPort")
	user := os.Getenv("pgUser")
	password := os.Getenv("pgPassword")
	dbname := os.Getenv("pgDbName")

	log.Println("host = ", host)
	log.Println("port = ", port)
	log.Println("user = ", user)
	log.Println("password = ", password)
	log.Println("dbname = ", dbname)

	log.Printf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
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
