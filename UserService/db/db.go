package db

import (
	"UserService/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Model:     gorm.Model{},
		Email:     "admin@gmail.com",
		Password:  "admin",
		FirstName: "Adminko",
		LastName:  "Adminic",
		Contact:   "231321",
		Role:      models.ADMIN,
		Banned:    false,
	},
	{
		Model:     gorm.Model{},
		Email:     "pera@gmail.com",
		Password:  "admin",
		FirstName: "Pera",
		LastName:  "Peric",
		Contact:   "231321",
		Role:      models.APPUSER,
		Banned:    false,
	},
	{
		Model:     gorm.Model{},
		Email:     "zika@gmail.com",
		Password:  "admin",
		FirstName: "Zika",
		LastName:  "Zikic",
		Contact:   "231321",
		Role:      models.DELIVERER,
		Banned:    false,
	},
	{
		Model:     gorm.Model{},
		Email:     "dunja@gmail.com",
		Password:  "admin",
		FirstName: "Dunja",
		LastName:  "Dunjica",
		Contact:   "231321",
		Role:      models.EMPLOYEE,
		Banned:    false,
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
