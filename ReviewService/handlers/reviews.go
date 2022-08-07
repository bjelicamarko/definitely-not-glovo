package handlers

import (
	"ReviewService/models"
	"ReviewService/repository"
	"ReviewService/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ReviewsHandler struct {
	repository *repository.Repository
}

func NewReviewsHandler(repository *repository.Repository) *ReviewsHandler {
	return &ReviewsHandler{repository}
}

func (rh *ReviewsHandler) GetReviewsOfRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	reviewsDTO, totalElements, _ := rh.repository.GetReviewsOfRestaurant(req)

	json.NewEncoder(resWriter).Encode(models.ReviewsPageable{Elements: reviewsDTO, TotalElements: totalElements})
}

func (rh *ReviewsHandler) SearchReviews(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	reviewsDTO, totalElements, _ := rh.repository.SearchReviews(req)

	json.NewEncoder(resWriter).Encode(models.ReviewsPageable{Elements: reviewsDTO, TotalElements: totalElements})
}

func (rh *ReviewsHandler) FindReviewByOrder(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	reviewDTO, err := rh.repository.FindReviewByOrder(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: "review successfully found", ReviewDTO: *reviewDTO})
}

func (rh *ReviewsHandler) CreateReview(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	var newReviewDTO models.ReviewDTO
	json.NewDecoder(req.Body).Decode(&newReviewDTO)

	reviewDTO, err := rh.repository.CreateReview(&newReviewDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: "review successfully created", ReviewDTO: *reviewDTO})
}

func (rh *ReviewsHandler) ReportReview(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	var updatedReviewDTO models.ReviewDTO
	json.NewDecoder(req.Body).Decode(&updatedReviewDTO)

	reviewDTO, err := rh.repository.ReportReview(&updatedReviewDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: "review successfully reported", ReviewDTO: *reviewDTO})
}

func (rh *ReviewsHandler) DeleteReview(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	reviewDTO, err := rh.repository.DeleteReview(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ReviewDTOMessage{Message: "review successfully deleted", ReviewDTO: *reviewDTO})

}
