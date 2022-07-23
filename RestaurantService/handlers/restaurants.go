package handlers

import (
	"RestaurantService/models"
	"RestaurantService/repository"
	"RestaurantService/utils"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type RestaurantsHandler struct {
	repository *repository.Repository
}

func NewRestaurantsHandler(repository *repository.Repository) *RestaurantsHandler {
	return &RestaurantsHandler{repository}
}

func (rh *RestaurantsHandler) GetRestaurants(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	restaurants, totalElements, _ := rh.repository.FindAll(req)

	json.NewEncoder(resWriter).Encode(models.RestaurantsPageable{Elements: restaurants, TotalElements: totalElements})
}

func (rh *RestaurantsHandler) CreateRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var newRestaurant models.RestaurantDTO
	json.NewDecoder(req.Body).Decode(&newRestaurant)

	restaurant, err := rh.repository.SaveRestaurant(&newRestaurant)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: err.Error()})
		return
	}

	_ = os.Remove(newRestaurant.ImagePath)
	utils.ToImage(newRestaurant.Image, newRestaurant.ImagePath)

	json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: "Restaurant successfully created", RestaurantDTO: restaurant.ToRestaurantDTO()})
}

func (rh *RestaurantsHandler) UpdateRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var updatedRestaurant models.Restaurant
	json.NewDecoder(req.Body).Decode(&updatedRestaurant)

	_, err := rh.repository.UpdateRestaurant(&updatedRestaurant)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: err.Error()})
	} else {
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Restaurant successfully updated"})
	}
}

func (rh *RestaurantsHandler) DeleteRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	err := rh.repository.DeleteRestaurant(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "error while deleting restaurant"})
	} else {
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Restaurant successfully deleted"})
	}
}
