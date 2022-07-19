package router

import (
	"RestaurantService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.RestaurantsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/restaurants/getRestaurants", handler.GetRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/createRestaurant", handler.CreateRestaurant).Methods(http.MethodPost)
	router.HandleFunc("/api/restaurants/updateRestaurant", handler.UpdateRestaurant).Methods(http.MethodPut)
	router.HandleFunc("/api/restaurants/deleteRestaurant/{id:[0-9]+}", handler.DeleteRestaurant).Methods(http.MethodDelete)

	http.ListenAndServe(":8082", router)
}
