package router

import (
	"UserService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/users/register", handler.Register).Methods("POST")

	router.HandleFunc("/api/users/authorize/admin", handler.AuthorizeAdmin).Methods("GET")
	router.HandleFunc("/api/users/authorize/appuser", handler.AuthorizeAppUser).Methods("GET")
	router.HandleFunc("/api/users/authorize/employee", handler.AuthorizeEmployee).Methods("GET")
	router.HandleFunc("/api/users/authorize/deliverer", handler.AuthorizeDeliverer).Methods("GET")

	http.ListenAndServe(":8081", router)
}
