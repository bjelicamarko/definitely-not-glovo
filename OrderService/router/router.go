package router

import (
	"OrderService/handlers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.OrdersHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"message": "Welcome to Dockerized Order Backend",
		}
		json.NewEncoder(rw).Encode(response)
	})

	router.HandleFunc("/api/orders/ordersForReport", handler.OrdersForReport).Methods("GET")
	router.HandleFunc("/api/orders/searchOrders", handler.SearchOrders).Methods("GET")
	router.HandleFunc("/api/orders/searchOrdersByRestaurantAndStatus", handler.SearchOrdersByRestaurantAndStatus).Methods("GET")
	router.HandleFunc("/api/orders/findOrderById/{id:[0-9]+}", handler.FindOrderById).Methods("GET")
	router.HandleFunc("/api/orders/reviewOrder/{id:[0-9]+}", handler.ReviewOrder).Methods("PATCH")

	router.HandleFunc("/api/orders/createOrder", handler.CreateOrder).Methods("POST")
	router.HandleFunc("/api/orders/changeStatusOfOrder", handler.ChangeStatusOfOrder).Methods("PUT")

	http.ListenAndServe(":8084", router)
}
