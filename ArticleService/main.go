package main

import (
	"ArticleService/db"
	"ArticleService/handlers"
	"ArticleService/repository"
	"ArticleService/router"
)

func main() {
	dbConn := db.Init()
	repository := repository.NewRepository(dbConn)
	articlesHandler := handlers.NewArticlesHandler(repository)
	router.MapRoutesAndServe(articlesHandler)
}
