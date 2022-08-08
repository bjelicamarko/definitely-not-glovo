package handlers

import (
	"GatewayService/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAllRestaurants(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.RestaurantsServiceRoot.Next().Host + RestaurantsServiceApi + "/findAllRestaurants?page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func SearchRestaurants(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	searchField := r.URL.Query().Get("searchField")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.RestaurantsServiceRoot.Next().Host + RestaurantsServiceApi + "/searchRestaurants?searchField=" + searchField + "&page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func FindRestaurantById(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	params := mux.Vars(r)
	restaurantId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(
		utils.RestaurantsServiceRoot.Next().Host + RestaurantsServiceApi + "/findRestaurantById/" + strconv.FormatUint(uint64(restaurantId), 10))

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func FindRestaurantByName(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	params := mux.Vars(r)
	restaurantName := params["name"]

	response, err := http.Get(
		utils.RestaurantsServiceRoot.Next().Host + RestaurantsServiceApi + "/findRestaurantByName/" + restaurantName)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func CreateRestaurant(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	req, _ := http.NewRequest(http.MethodPost,
		utils.RestaurantsServiceRoot.Next().Host+RestaurantsServiceApi+"/createRestaurant", r.Body)
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

func UpdateRestaurant(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	req, _ := http.NewRequest(http.MethodPut,
		utils.RestaurantsServiceRoot.Next().Host+RestaurantsServiceApi+"/updateRestaurant", r.Body)
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

func DeleteRestaurant(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete,
		utils.RestaurantsServiceRoot.Next().Host+RestaurantsServiceApi+"/deleteRestaurant/"+strconv.FormatUint(uint64(userId), 10),
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
