package handlers

import (
	"GatewayService/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func SearchOrders(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	role := r.URL.Query().Get("role")
	userId := r.URL.Query().Get("userId")
	restaurantId := r.URL.Query().Get("restaurantId")
	orderStatus := r.URL.Query().Get("orderStatus")
	priceFrom := r.URL.Query().Get("priceFrom")
	priceTo := r.URL.Query().Get("priceTo")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.OrdersServiceRoot.Next().Host + OrdersServiceApi +
			"/searchOrders?role=" + role + "&userId=" + userId + "&restaurantId=" + restaurantId +
			"&orderStatus=" + orderStatus + "&priceFrom=" + priceFrom +
			"&priceTo=" + priceTo + "&page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func SearchOrdersByRestaurantAndStatus(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	restaurantId := r.URL.Query().Get("restaurantId")
	orderStatus := r.URL.Query().Get("orderStatus")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.OrdersServiceRoot.Next().Host + OrdersServiceApi +
			"/searchOrdersByRestaurantAndStatus?restaurantId=" + restaurantId +
			"&orderStatus=" + orderStatus +
			"&page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func FindOrderById(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	orderId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(
		utils.OrdersServiceRoot.Next().Host + OrdersServiceApi + "/findOrderById/" + strconv.FormatUint(uint64(orderId), 10))

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func ReviewOrder(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "appuser") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	orderId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPatch,
		utils.OrdersServiceRoot.Next().Host+OrdersServiceApi+"/reviewOrder/"+strconv.FormatUint(uint64(orderId), 10),
		r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func CreateOrder(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "appuser") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPost,
		utils.OrdersServiceRoot.Next().Host+OrdersServiceApi+"/createOrder", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func ChangeStatusOfOrder(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "appuser") != nil && utils.AuthorizeRole(r, "deliverer") != nil &&
		utils.AuthorizeRole(r, "employee") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPut,
		utils.OrdersServiceRoot.Next().Host+OrdersServiceApi+"/changeStatusOfOrder", r.Body)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}
