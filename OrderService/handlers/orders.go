package handlers

import (
	"OrderService/models"
	"OrderService/repository"
	"OrderService/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrdersHandler struct {
	repository *repository.Repository
}

func NewOrdersHandler(repository *repository.Repository) *OrdersHandler {
	return &OrdersHandler{repository}
}

func (oh *OrdersHandler) SearchOrdersByRestaurantAndStatus(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	ordersDTO, totalElements, _ := oh.repository.SearchOrdersByRestaurantAndStatus(req)

	json.NewEncoder(resWriter).Encode(models.OrdersPageable{Elements: ordersDTO, TotalElements: totalElements})
}

func (oh *OrdersHandler) FindOrderById(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	orderDTO, err := oh.repository.FindOrderById(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.OrderDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.OrderDTOMessage{Message: "order successfully found", OrderDTO: *orderDTO})
}

func (oh *OrdersHandler) CreateOrder(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	var newOrderDTO models.OrderDTO
	json.NewDecoder(req.Body).Decode(&newOrderDTO)

	orderDTO, err := oh.repository.CreateOrder(&newOrderDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.OrderDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.OrderDTOMessage{Message: "order successfully created", OrderDTO: *orderDTO})
}

func (oh *OrdersHandler) ChangeStatusOfOrder(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	var orderStatusDTO models.OrderStatusDTO
	json.NewDecoder(req.Body).Decode(&orderStatusDTO)

	orderDTO, err := oh.repository.ChangeStatusOfOrder(&orderStatusDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.OrderDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.OrderDTOMessage{Message: "order successfully changed", OrderDTO: *orderDTO})
}
