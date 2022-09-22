package handlers

import (
	"GatewayService/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetReviewsOfRestaurant(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	restaurantId := r.URL.Query().Get("restaurantId")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.ReviewsServiceRoot.Next().Host + ReviewsServiceApi + "/getReviewsOfRestaurant?restaurantId=" + restaurantId +
			"&page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func SearchReviews(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	restaurantId := r.URL.Query().Get("restaurantId")
	userId := r.URL.Query().Get("userId")
	inappropriate := r.URL.Query().Get("inappropriate")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.ReviewsServiceRoot.Next().Host + ReviewsServiceApi + "/searchReviews?restaurantId=" + restaurantId +
			"&userId=" + userId + "&inappropriate=" + inappropriate +
			"&page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func FindReviewByOrder(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	orderId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(
		utils.ReviewsServiceRoot.Next().Host + ReviewsServiceApi + "/findReviewByOrder/" + strconv.FormatUint(uint64(orderId), 10))

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func CreateReview(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "appuser") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPost,
		utils.ReviewsServiceRoot.Next().Host+ReviewsServiceApi+"/createReview", r.Body)
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

func ReportReview(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "employee") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPut,
		utils.ReviewsServiceRoot.Next().Host+ReviewsServiceApi+"/reportReview", r.Body)
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

func DeleteReview(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete,
		utils.ReviewsServiceRoot.Next().Host+ReviewsServiceApi+"/deleteReview/"+strconv.FormatUint(uint64(reviewId), 10),
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

func AverageRatingOfRestaurant(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	restaurantId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(
		utils.ReviewsServiceRoot.Next().Host + ReviewsServiceApi + "/averageRatingOfRestaurant/" + strconv.FormatUint(uint64(restaurantId), 10))

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}
