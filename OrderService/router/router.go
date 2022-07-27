package router

import (
	"OrderService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.OrdersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/orders/searchOrdersByRestaurantAndStatus", handler.SearchOrdersByRestaurantAndStatus).Methods("GET")
	router.HandleFunc("/api/orders/findOrderById/{id:[0-9]+}", handler.FindOrderById).Methods("GET")
	router.HandleFunc("/api/orders/createOrder", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/changeStatusOfOrder", handler.ChangeStatusOfOrder).Methods("PUT")

	http.ListenAndServe(":8084", router)
}
