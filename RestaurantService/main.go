package main

import (
	"RestaurantService/db"
	"RestaurantService/handlers"
	"RestaurantService/repository"
	"RestaurantService/router"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	restaurantHandler := handlers.NewRestaurantsHandler(repository)
	router.MapRoutesAndServe(restaurantHandler)
}
