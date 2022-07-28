package main

import (
	"RestaurantService/db"
	"RestaurantService/handlers"
	"RestaurantService/repository"
	"RestaurantService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	restaurantHandler := handlers.NewRestaurantsHandler(repository)
	router.MapRoutesAndServe(restaurantHandler)
}
