package db

import (
	"OrderService/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var orders = []models.Order{
	{
		Model:          gorm.Model{},
		IdRestaurant:   1,
		RestaurantName: "McDonalds",
		IdAppUser:      2,
		IdEmployee:     4,
		IdDeliverer:    3,
		OrderStatus:    models.DELIVERED,
		TotalPrice:     1400.0,
		Tip:            50.0,
		Note:           "porudzbina 1",
		DateTime:       "17.07.2022. 17:00",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Luka Petrovic",
		StreetNumber:   "4",
		DisplayName:    "Trebinje Luka Petrovic 4",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
		Reviewed:       true,
	},
	{
		Model:          gorm.Model{},
		IdRestaurant:   2,
		RestaurantName: "KFC",
		IdAppUser:      2,
		IdEmployee:     6,
		IdDeliverer:    7,
		OrderStatus:    models.DELIVERED,
		TotalPrice:     2200.0,
		Tip:            0.0,
		Note:           "porudzbina 2",
		DateTime:       "18.07.2022. 18:00",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Nikola Tesla",
		StreetNumber:   "3",
		DisplayName:    "Trebinje Nikola Tesla 3",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
		Reviewed:       true,
	},
	{
		Model:          gorm.Model{},
		IdRestaurant:   3,
		RestaurantName: "FOODIC",
		IdAppUser:      2,
		IdEmployee:     9,
		IdDeliverer:    10,
		OrderStatus:    models.DELIVERED,
		TotalPrice:     1280.0,
		Tip:            100.0,
		Note:           "porudzbina 3",
		DateTime:       "19.07.2022. 19:00",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Nikola Tesla",
		StreetNumber:   "3",
		DisplayName:    "Trebinje Nikola Tesla 3",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
		Reviewed:       true,
	},
	{
		Model:          gorm.Model{},
		IdRestaurant:   1,
		RestaurantName: "McDonalds",
		IdAppUser:      2,
		IdEmployee:     4,
		IdDeliverer:    3,
		OrderStatus:    models.DELIVERED,
		TotalPrice:     1600.0,
		Tip:            100.0,
		Note:           "porudzbina 4",
		DateTime:       "20.07.2022. 20:00",
		Country:        "Bosnia and Hercegovina",
		City:           "Trebinje",
		Street:         "Vojvoda Stepa",
		StreetNumber:   "10",
		DisplayName:    "Trebinje Vojvoda Stepa 10",
		Ptt:            25000,
		Longitude:      18.3501358,
		Latitude:       42.7060377,
		Reviewed:       true,
	},
}

var orderItems = []models.OrderItem{
	{
		Model:        gorm.Model{},
		IdOrder:      1,
		IdArticle:    1,
		ArticleName:  "Punjeno bijelo meso",
		CurrentPrice: 400.0,
		Quantity:     2,
		TotalPrice:   800.0,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      1,
		IdArticle:    2,
		ArticleName:  "Coca-Cola 0.5",
		CurrentPrice: 200,
		Quantity:     2,
		TotalPrice:   400,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      2,
		IdArticle:    4,
		ArticleName:  "Water 0.5",
		CurrentPrice: 150.0,
		Quantity:     3,
		TotalPrice:   450.0,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      2,
		IdArticle:    5,
		ArticleName:  "Chicken Nuggets",
		CurrentPrice: 500.0,
		Quantity:     3,
		TotalPrice:   1500.0,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      3,
		IdArticle:    9,
		ArticleName:  "Index sandwich",
		CurrentPrice: 350.0,
		Quantity:     2,
		TotalPrice:   700.0,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      3,
		IdArticle:    8,
		ArticleName:  "Sprite 0.5",
		CurrentPrice: 200.0,
		Quantity:     2,
		TotalPrice:   400.0,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      4,
		IdArticle:    3,
		ArticleName:  "Big Mac",
		CurrentPrice: 500.0,
		Quantity:     2,
		TotalPrice:   1000.0,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      4,
		IdArticle:    2,
		ArticleName:  "Coca-Cola 0.5",
		CurrentPrice: 200.0,
		Quantity:     2,
		TotalPrice:   400,
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

	db.Migrator().DropTable("order_items")
	db.Migrator().DropTable("orders")

	db.Migrator().AutoMigrate(&models.Order{})
	db.Migrator().AutoMigrate(&models.OrderItem{})

	for _, order := range orders {
		db.Create(&order)
	}

	for _, orderItem := range orderItems {
		db.Create(&orderItem)
	}

	return db
}
