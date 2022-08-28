package router

import (
	"ArticleService/handlers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.ArticlesHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Welcome to Dockerized Articles Backend",
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/api/articles/findAllArticles", handler.FindAllArticles).Methods(http.MethodGet)
	router.HandleFunc("/api/articles/findAllArticlesFromRestaurant", handler.FindAllArticlesFromRestaurant).Methods(http.MethodGet)
	router.HandleFunc("/api/articles/searchArticles", handler.SearchArticles).Methods(http.MethodGet)
	router.HandleFunc("/api/articles/findArticleById/{id:[0-9]+}", handler.FindArticleById).Methods(http.MethodGet)

	router.HandleFunc("/api/articles/createArticle", handler.CreateArticle).Methods(http.MethodPost)
	router.HandleFunc("/api/articles/updateArticle", handler.UpdateArticle).Methods(http.MethodPut)
	router.HandleFunc("/api/articles/deleteArticle/{id:[0-9]+}", handler.DeleteArticle).Methods(http.MethodDelete)

	log.Println("Server is running!")

	http.ListenAndServe(":8083", router)

}
