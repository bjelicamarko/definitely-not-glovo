package main

import (
	"OrderService/db"
	"OrderService/handlers"
	"OrderService/repository"
	"OrderService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	ordersHandler := handlers.NewOrdersHandler(repository)
	router.MapRoutesAndServe(ordersHandler)
}
