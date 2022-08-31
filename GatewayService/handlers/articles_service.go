package handlers

import (
	"GatewayService/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindAllArticles(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.ArticlesServiceRoot.Next().Host + ArticlesServiceApi + "/findAllArticles?page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func FindAllArticlesFromRestaurant(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	restaurantName := r.URL.Query().Get("restaurantName")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.ArticlesServiceRoot.Next().Host + ArticlesServiceApi + "/findAllArticlesFromRestaurant?restaurantName=" + restaurantName + "&page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func SearchArticles(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	restaurantName := r.URL.Query().Get("restaurantName")
	searchField := r.URL.Query().Get("searchField")
	articleType := r.URL.Query().Get("articleType")
	priceFrom := r.URL.Query().Get("priceFrom")
	priceTo := r.URL.Query().Get("priceTo")
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(
		utils.ArticlesServiceRoot.Next().Host + ArticlesServiceApi + "/searchArticles?restaurantName=" + restaurantName +
			"&searchField=" + searchField + "&articleType=" + articleType +
			"&priceFrom=" + priceFrom + "&priceTo=" + priceTo +
			"&page=" + page + "&size=" + size)

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func FindArticleById(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	restaurantId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(
		utils.ArticlesServiceRoot.Next().Host + ArticlesServiceApi + "/findArticleById/" + strconv.FormatUint(uint64(restaurantId), 10))

	if err != nil {
		resWriter.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, resWriter)
}

func CreateArticle(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPost,
		utils.ArticlesServiceRoot.Next().Host+ArticlesServiceApi+"/createArticle", r.Body)
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

func UpdateArticle(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPut,
		utils.ArticlesServiceRoot.Next().Host+ArticlesServiceApi+"/updateArticle", r.Body)
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

func DeleteArticle(resWriter http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&resWriter, r)

	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete,
		utils.ArticlesServiceRoot.Next().Host+ArticlesServiceApi+"/deleteArticle/"+strconv.FormatUint(uint64(userId), 10),
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
