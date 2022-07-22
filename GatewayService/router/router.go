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
	router.HandleFunc("/api/users/getUsers", handlers.FindAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/searchUsers", handlers.SeachUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/updateUser", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/deleteUser/{id:[0-9]+}", handlers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/banUser/{id:[0-9]+}", handlers.BanUser).Methods(http.MethodPatch)
	router.HandleFunc("/api/users/unbanUser/{id:[0-9]+}", handlers.UnbanUser).Methods(http.MethodPatch)

	router.HandleFunc("/api/restaurants/getRestaurants", handlers.FindAllRestaurants).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
