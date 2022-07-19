package router

import (
	"GatewayService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handlers.Login).Methods(http.MethodPost)

	router.HandleFunc("/api/restaurants/getRestaurants", handlers.FindAllRestaurants).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
