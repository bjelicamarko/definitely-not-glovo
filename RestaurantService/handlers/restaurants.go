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

func (rh *RestaurantsHandler) FindAllRestaurants(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	restaurantsDTO, totalElements, _ := rh.repository.FindAllRestaurants(req)

	json.NewEncoder(resWriter).Encode(models.RestaurantsPageable{Elements: restaurantsDTO, TotalElements: totalElements})
}

func (rh *RestaurantsHandler) SearchRestaurants(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	restaurantsDTO, totalElements, _ := rh.repository.SearchRestaurants(req)

	json.NewEncoder(resWriter).Encode(models.RestaurantsPageable{Elements: restaurantsDTO, TotalElements: totalElements})
}

func (rh *RestaurantsHandler) FindRestaurantById(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	restaurantDTO, err := rh.repository.FindRestaurantById(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: "restaurant successfully found", RestaurantDTO: *restaurantDTO})
}

func (rh *RestaurantsHandler) CreateRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var newRestaurantDTO models.RestaurantDTO
	json.NewDecoder(req.Body).Decode(&newRestaurantDTO)

	_ = os.Remove(newRestaurantDTO.ImagePath)
	utils.ToImage(newRestaurantDTO.Image, newRestaurantDTO.ImagePath)

	restaurantDTO, err := rh.repository.CreateRestaurant(&newRestaurantDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: "restaurant successfully created", RestaurantDTO: *restaurantDTO})
}

func (rh *RestaurantsHandler) UpdateRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var updatedRestaurantDTO models.RestaurantDTO
	json.NewDecoder(req.Body).Decode(&updatedRestaurantDTO)

	if updatedRestaurantDTO.Changed {
		_ = os.Remove(updatedRestaurantDTO.ImagePath)
		utils.ToImage(updatedRestaurantDTO.Image, updatedRestaurantDTO.ImagePath)
	}

	restaurantDTO, err := rh.repository.UpdateRestaurant(&updatedRestaurantDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: "restaurant successfully updated", RestaurantDTO: *restaurantDTO})
}

func (rh *RestaurantsHandler) DeleteRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	restaurantDTO, err := rh.repository.DeleteRestaurant(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.RestaurantDTOMessage{Message: "restaurant successfully deleted", RestaurantDTO: *restaurantDTO})
}
