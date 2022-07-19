package main

import (
	"UserService/db"
	"UserService/handlers"
	"UserService/repository"
	"UserService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	usersHandler := handlers.NewUsersHandler(repository)
	router.MapRoutesAndServe(usersHandler)
}
