package router

import (
	"GatewayService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/register", handlers.Register).Methods(http.MethodPost)

	router.HandleFunc("/api/users/findAllUsers", handlers.FindAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/searchUsers", handlers.SeachUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/findUserById/{id:[0-9]+}", handlers.FindUserById).Methods(http.MethodGet)
	router.HandleFunc("/api/users/createUser", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/updateUser", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/deleteUser/{id:[0-9]+}", handlers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/banUser/{id:[0-9]+}", handlers.BanUser).Methods(http.MethodPatch)
	router.HandleFunc("/api/users/unbanUser/{id:[0-9]+}", handlers.UnbanUser).Methods(http.MethodPatch)

	router.HandleFunc("/api/restaurants/findAllRestaurants", handlers.FindAllRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/searchRestaurants", handlers.SearchRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/findRestaurantById/{id:[0-9]+}", handlers.FindRestaurantById).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/createRestaurant", handlers.CreateRestaurant).Methods(http.MethodPost)
	router.HandleFunc("/api/restaurants/updateRestaurant", handlers.UpdateRestaurant).Methods(http.MethodPut)
	router.HandleFunc("/api/restaurants/deleteRestaurant/{id:[0-9]+}", handlers.DeleteRestaurant).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)
}
