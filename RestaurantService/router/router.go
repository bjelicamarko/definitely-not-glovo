package router

import (
	"RestaurantService/handlers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.RestaurantsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Welcome to Dockerized Restaurants Backend",
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/api/restaurants/findAllRestaurants", handler.FindAllRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/searchRestaurants", handler.SearchRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/findRestaurantById/{id:[0-9]+}", handler.FindRestaurantById).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/findRestaurantByName/{name}", handler.FindRestaurantByName).Methods(http.MethodGet)

	router.HandleFunc("/api/restaurants/createRestaurant", handler.CreateRestaurant).Methods(http.MethodPost)
	router.HandleFunc("/api/restaurants/updateRestaurant", handler.UpdateRestaurant).Methods(http.MethodPut)
	router.HandleFunc("/api/restaurants/deleteRestaurant/{id:[0-9]+}", handler.DeleteRestaurant).Methods(http.MethodDelete)

	http.ListenAndServe(":8082", router)
}
