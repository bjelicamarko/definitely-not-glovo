package router

import (
	"UserService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/login", handler.Login).Methods("POST")

	router.HandleFunc("/authorize/admin", handler.AuthorizeAdmin).Methods("GET")
	router.HandleFunc("/authorize/appuser", handler.AuthorizeAppUser).Methods("GET")
	router.HandleFunc("/authorize/employee", handler.AuthorizeEmployee).Methods("GET")
	router.HandleFunc("/authorize/deliverer", handler.AuthorizeDeliverer).Methods("GET")

	http.ListenAndServe(":8081", router)
}
