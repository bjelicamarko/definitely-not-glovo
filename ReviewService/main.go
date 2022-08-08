package main

import (
	"ReviewService/db"
	"ReviewService/handlers"
	"ReviewService/repository"
	"ReviewService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	reviewsHandler := handlers.NewReviewsHandler(repository)
	router.MapRoutesAndServe(reviewsHandler)
}
