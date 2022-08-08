package router

import (
	"ReviewService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe(handler *handlers.ReviewsHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/api/reviews/getReviewsOfRestaurant", handler.GetReviewsOfRestaurant).Methods(http.MethodGet)
	router.HandleFunc("/api/reviews/searchReviews", handler.SearchReviews).Methods(http.MethodGet)
	router.HandleFunc("/api/reviews/findReviewByOrder/{id:[0-9]+}", handler.FindReviewByOrder).Methods(http.MethodGet)

	router.HandleFunc("/api/reviews/createReview", handler.CreateReview).Methods(http.MethodPost)
	router.HandleFunc("/api/reviews/reportReview", handler.ReportReview).Methods(http.MethodPut)
	router.HandleFunc("/api/reviews/deleteReview/{id:[0-9]+}", handler.DeleteReview).Methods(http.MethodDelete)

	router.HandleFunc("/api/reviews/averageRatingOfRestaurant/{id:[0-9]+}", handler.AverageRatingOfRestaurant).Methods(http.MethodGet)

	http.ListenAndServe(":8085", router)
}
