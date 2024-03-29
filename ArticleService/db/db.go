package db

import (
	"ArticleService/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var articles = []models.Article{
	{
		Model:          gorm.Model{}, // 1
		ArticleName:    "Punjeno bijelo meso",
		ArticleType:    "FOOD",
		Price:          400.0,
		Description:    "Tasty",
		RestaurantName: "McDonalds",
		Image:          "images/default.jpg",
	},
	{
		Model:          gorm.Model{}, // 2
		ArticleName:    "Coca-Cola 0.5",
		ArticleType:    "DRINK",
		Price:          200.0,
		Description:    "Refreshing",
		RestaurantName: "McDonalds",
		Image:          "images/drinks.jpg",
	},
	{
		Model:          gorm.Model{}, // 3
		ArticleName:    "Big Mac",
		ArticleType:    "FOOD",
		Price:          500.0,
		Description:    "Tasty",
		RestaurantName: "McDonalds",
		Image:          "images/default.jpg",
	},
	{
		Model:          gorm.Model{}, // 4
		ArticleName:    "Water 0.5",
		ArticleType:    "DRINK",
		Price:          150.0,
		Description:    "Refreshing",
		RestaurantName: "KFC",
		Image:          "images/drinks.jpg",
	},
	{
		Model:          gorm.Model{}, // 5
		ArticleName:    "Chicken Nuggets",
		ArticleType:    "FOOD",
		Price:          500.0,
		Description:    "Tasty",
		RestaurantName: "KFC",
		Image:          "images/default.jpg",
	},
	{
		Model:          gorm.Model{}, // 6
		ArticleName:    "Coca Cola Zero 0.5",
		ArticleType:    "DRINK",
		Price:          200.0,
		Description:    "Refreshing",
		RestaurantName: "KFC",
		Image:          "images/drinks.jpg",
	},
	{
		Model:          gorm.Model{}, // 7
		ArticleName:    "Premium piletina sandwich",
		ArticleType:    "FOOD",
		Price:          450.0,
		Description:    "Tasty",
		RestaurantName: "FOODIC",
		Image:          "images/default.jpg",
	},
	{
		Model:          gorm.Model{}, // 8
		ArticleName:    "Sprite 0.5",
		ArticleType:    "DRINK",
		Price:          200.0,
		Description:    "Refreshing",
		RestaurantName: "FOODIC",
		Image:          "images/drinks.jpg",
	},
	{
		Model:          gorm.Model{}, // 9
		ArticleName:    "Index sandwich",
		ArticleType:    "FOOD",
		Price:          350.0,
		Description:    "Tasty",
		RestaurantName: "FOODIC",
		Image:          "images/default.jpg",
	},
	{
		Model:          gorm.Model{}, // 10
		ArticleName:    "Orange Juice 0.3",
		ArticleType:    "DRINK",
		Price:          180.0,
		Description:    "Refreshing",
		RestaurantName: "FOODIC",
		Image:          "images/drinks.jpg",
	},
	{
		Model:          gorm.Model{}, // 11
		ArticleName:    "Chips",
		ArticleType:    "FOOD",
		Price:          200.0,
		Description:    "Tasty",
		RestaurantName: "FOODIC",
		Image:          "images/default.jpg",
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

	db.Migrator().DropTable("articles")
	db.Migrator().AutoMigrate(&models.Article{})

	for _, article := range articles {
		db.Create(&article)
	}

	return db
}
