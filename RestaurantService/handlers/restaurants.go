package handlers

import (
	"RestaurantService/models"
	"RestaurantService/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type RestaurantsHandler struct {
	repository *repository.Repository
}

func NewRestaurantsHandler(repository *repository.Repository) *RestaurantsHandler {
	return &RestaurantsHandler{repository}
}

func (rh *RestaurantsHandler) GetRestaurants(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	restaurants, _ := rh.repository.FindAll()

	json.NewEncoder(resWriter).Encode(restaurants)
}

func (rh *RestaurantsHandler) CreateRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var newRestaurant models.Restaurant
	json.NewDecoder(req.Body).Decode(&newRestaurant)

	newRestaurant.Model = gorm.Model{}
	_, err := rh.repository.SaveRestaurant(&newRestaurant)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: err.Error()})
	} else {
		json.NewEncoder(resWriter).Encode(models.Response{Message: "Restaurant successfully created"})
	}
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
