package router

import (
	"GatewayService/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func MapRoutesAndServe() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users/login", handlers.Login).Methods(http.MethodPost)
	router.HandleFunc("/api/users/register", handlers.Register).Methods(http.MethodPost)

	router.HandleFunc("/api/users/findAllUsers", handlers.FindAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/searchUsers", handlers.SeachUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/users/findUserById/{id:[0-9]+}", handlers.FindUserById).Methods(http.MethodGet)
	router.HandleFunc("/api/users/createUser", handlers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/users/updateUser", handlers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/api/users/deleteUser/{id:[0-9]+}", handlers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/users/banUser/{id:[0-9]+}", handlers.BanUser).Methods(http.MethodPatch)
	router.HandleFunc("/api/users/unbanUser/{id:[0-9]+}", handlers.UnbanUser).Methods(http.MethodPatch)

	router.HandleFunc("/api/restaurants/findAllRestaurants", handlers.FindAllRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/searchRestaurants", handlers.SearchRestaurants).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/findRestaurantById/{id:[0-9]+}", handlers.FindRestaurantById).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/findRestaurantByName/{name}", handlers.FindRestaurantByName).Methods(http.MethodGet)
	router.HandleFunc("/api/restaurants/createRestaurant", handlers.CreateRestaurant).Methods(http.MethodPost)
	router.HandleFunc("/api/restaurants/updateRestaurant", handlers.UpdateRestaurant).Methods(http.MethodPut)
	router.HandleFunc("/api/restaurants/deleteRestaurant/{id:[0-9]+}", handlers.DeleteRestaurant).Methods(http.MethodDelete)

	router.HandleFunc("/api/articles/findAllArticles", handlers.FindAllArticles).Methods(http.MethodGet)
	router.HandleFunc("/api/articles/findAllArticlesFromRestaurant", handlers.FindAllArticlesFromRestaurant).Methods(http.MethodGet)
	router.HandleFunc("/api/articles/searchArticles", handlers.SearchArticles).Methods(http.MethodGet)
	router.HandleFunc("/api/articles/findArticleById/{id:[0-9]+}", handlers.FindArticleById).Methods(http.MethodGet)
	router.HandleFunc("/api/articles/createArticle", handlers.CreateArticle).Methods(http.MethodPost)
	router.HandleFunc("/api/articles/updateArticle", handlers.UpdateArticle).Methods(http.MethodPut)
	router.HandleFunc("/api/articles/deleteArticle/{id:[0-9]+}", handlers.DeleteArticle).Methods(http.MethodDelete)

	router.HandleFunc("/api/orders/searchOrders", handlers.SearchOrders).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/searchOrdersByRestaurantAndStatus", handlers.SearchOrdersByRestaurantAndStatus).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/findOrderById/{id:[0-9]+}", handlers.FindOrderById).Methods(http.MethodGet)
	router.HandleFunc("/api/orders/createOrder", handlers.CreateOrder).Methods(http.MethodPost)
	router.HandleFunc("/api/orders/changeStatusOfOrder", handlers.ChangeStatusOfOrder).Methods(http.MethodPut)

	router.HandleFunc("/api/reviews/getReviewsOfRestaurant", handlers.GetReviewsOfRestaurant).Methods(http.MethodGet)
	router.HandleFunc("/api/reviews/searchReviews", handlers.SearchReviews).Methods(http.MethodGet)
	router.HandleFunc("/api/reviews/findReviewByOrder/{id:[0-9]+}", handlers.FindReviewByOrder).Methods(http.MethodGet)
	router.HandleFunc("/api/reviews/createReview", handlers.CreateReview).Methods(http.MethodPost)
	router.HandleFunc("/api/reviews/reportReview", handlers.ReportReview).Methods(http.MethodPut)
	router.HandleFunc("/api/reviews/deleteReview/{id:[0-9]+}", handlers.DeleteReview).Methods(http.MethodDelete)
	router.HandleFunc("/api/reviews/averageRatingOfRestaurant/{id:[0-9]+}", handlers.AverageRatingOfRestaurant).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
}
