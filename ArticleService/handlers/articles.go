package handlers

import (
	"ArticleService/models"
	"ArticleService/repository"
	"ArticleService/utils"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type ArticlesHandler struct {
	repository *repository.Repository
}

func NewArticlesHandler(repository *repository.Repository) *ArticlesHandler {
	return &ArticlesHandler{repository}
}

func (ah *ArticlesHandler) FindAllArticles(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	articlesDTO, totalElements, _ := ah.repository.FindAllArticles(req)

	json.NewEncoder(resWriter).Encode(models.ArticlesPageable{Elements: articlesDTO, TotalElements: totalElements})
}

func (ah *ArticlesHandler) FindAllArticlesFromRestaurant(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	articlesDTO, totalElements, _ := ah.repository.FindAllArticlesFromRestaurant(req)

	json.NewEncoder(resWriter).Encode(models.ArticlesPageable{Elements: articlesDTO, TotalElements: totalElements})
}

func (ah *ArticlesHandler) SearchArticles(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	articlesDTO, totalElements, _ := ah.repository.SearchArticles(req)

	json.NewEncoder(resWriter).Encode(models.ArticlesPageable{Elements: articlesDTO, TotalElements: totalElements})
}

func (ah *ArticlesHandler) FindArticleById(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	articleDTO, err := ah.repository.FindArticleById(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: "article successfully found", ArticleDTO: *articleDTO})
}

func (ah *ArticlesHandler) CreateArticle(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var newArticleDTO models.ArticleDTO
	json.NewDecoder(req.Body).Decode(&newArticleDTO)

	_ = os.Remove(newArticleDTO.ImagePath)
	utils.ToImage(newArticleDTO.Image, newArticleDTO.ImagePath)

	articleDTO, err := ah.repository.CreateArticle(&newArticleDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: "article successfully created", ArticleDTO: *articleDTO})
}

func (ah *ArticlesHandler) UpdateArticle(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var updatedArticleDTO models.ArticleDTO
	json.NewDecoder(req.Body).Decode(&updatedArticleDTO)

	if updatedArticleDTO.Changed {
		_ = os.Remove(updatedArticleDTO.ImagePath)
		utils.ToImage(updatedArticleDTO.Image, updatedArticleDTO.ImagePath)
	}

	articleDTO, err := ah.repository.UpdateArticle(&updatedArticleDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: "article successfully updated", ArticleDTO: *articleDTO})
}

func (ah *ArticlesHandler) DeleteArticle(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	articleDTO, err := ah.repository.DeleteArticle(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.ArticleDTOMessage{Message: "article successfully deleted", ArticleDTO: *articleDTO})
}
