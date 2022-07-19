package router

import (
	"RestaurantService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.RestaurantsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.GetRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/createRestaurant", handler.CreateRestaurant).Methods(http.MethodPost)
	router.HandleFunc("/updateRestaurant", handler.UpdateRestaurant).Methods(http.MethodPut)
	router.HandleFunc("/deleteRestaurant/{id:[0-9]+}", handler.DeleteRestaurant).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)
}
