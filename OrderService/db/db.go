package db

import (
	"OrderService/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var orders = []models.Order{
	{
		Model:        gorm.Model{},
		IdRestaurant: 1,
		IdAppUser:    2,
		IdEmployee:   4,
		IdDeliverer:  3,
		OrderStatus:  models.DELIVERED,
		TotalPrice:   1200.0,
		Tip:          50.0,
		Note:         "porudzbina 1",
		DateTime:     "27.7.2022. 17:00",
	},
	{
		Model:        gorm.Model{},
		IdRestaurant: 2,
		IdAppUser:    2,
		IdEmployee:   4,
		IdDeliverer:  3,
		OrderStatus:  models.DELIVERED,
		TotalPrice:   600.0,
		Tip:          0.0,
		Note:         "porudzbina 2",
		DateTime:     "27.7.2022. 18:00",
	},
	{
		Model:        gorm.Model{},
		IdRestaurant: 3,
		IdAppUser:    2,
		IdEmployee:   4,
		IdDeliverer:  3,
		OrderStatus:  models.DELIVERED,
		TotalPrice:   2400.0,
		Tip:          100.0,
		Note:         "porudzbina 3",
		DateTime:     "27.7.2022. 19:00",
	},
	{
		Model:        gorm.Model{},
		IdRestaurant: 1,
		IdAppUser:    2,
		IdEmployee:   4,
		IdDeliverer:  3,
		OrderStatus:  models.ACCEPTED,
		TotalPrice:   2400.0,
		Tip:          100.0,
		Note:         "porudzbina 4",
		DateTime:     "27.7.2022. 20:00",
	},
}

var orderItems = []models.OrderItem{
	{
		Model:        gorm.Model{},
		IdOrder:      1,
		IdArticle:    1,
		ArticleName:  "Coca Cola",
		CurrentPrice: 200,
		Quantity:     2,
		TotalPrice:   400,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      1,
		IdArticle:    2,
		ArticleName:  "Pizza",
		CurrentPrice: 400,
		Quantity:     2,
		TotalPrice:   800,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      2,
		IdArticle:    1,
		ArticleName:  "Coca Cola",
		CurrentPrice: 200,
		Quantity:     1,
		TotalPrice:   200,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      2,
		IdArticle:    3,
		ArticleName:  "Index",
		CurrentPrice: 400,
		Quantity:     1,
		TotalPrice:   400,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      3,
		IdArticle:    1,
		ArticleName:  "Coca Cola",
		CurrentPrice: 200,
		Quantity:     4,
		TotalPrice:   800,
	},
	{
		Model:        gorm.Model{},
		IdOrder:      3,
		IdArticle:    3,
		ArticleName:  "Index",
		CurrentPrice: 400,
		Quantity:     4,
		TotalPrice:   1600,
	},
}

func Init() *gorm.DB {
	dsn := "host=localhost user=postgres password=admin dbname=ordersDB port=5432 sslmode=disable"
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
