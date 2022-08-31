package router

import (
	"UserService/handlers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.UsersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Welcome to Dockerized Users Backend",
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/api/users/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/users/register", handler.Register).Methods("POST")

	router.HandleFunc("/api/users/authorize/admin", handler.AuthorizeAdmin).Methods("GET")
	router.HandleFunc("/api/users/authorize/appuser", handler.AuthorizeAppUser).Methods("GET")
	router.HandleFunc("/api/users/authorize/employee", handler.AuthorizeEmployee).Methods("GET")
	router.HandleFunc("/api/users/authorize/deliverer", handler.AuthorizeDeliverer).Methods("GET")

	router.HandleFunc("/api/users/findAllUsers", handler.FindAllUsers).Methods("GET")
	router.HandleFunc("/api/users/searchUsers", handler.SearchUsers).Methods("GET")
	router.HandleFunc("/api/users/findUserById/{id:[0-9]+}", handler.FindUserById).Methods("GET")

	router.HandleFunc("/api/users/createUser", handler.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/updateUser", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/deleteUser/{id:[0-9]+}", handler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/users/banUser/{id:[0-9]+}", handler.BanUser).Methods("PATCH")
	router.HandleFunc("/api/users/unbanUser/{id:[0-9]+}", handler.UnbanUser).Methods("PATCH")

	http.ListenAndServe(":8081", router)
}
