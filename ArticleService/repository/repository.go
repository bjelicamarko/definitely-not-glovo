package repository

import (
	"ArticleService/models"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func concat(str string) string {
	return "%" + strings.ToLower(str) + "%"
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page < 0 {
			page = 0
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (repo *Repository) FindAllArticles(r *http.Request) ([]models.ArticleDTO, int64, error) {
	var articlesDTO []models.ArticleDTO
	var articles []*models.Article
	var totalElements int64

	result := repo.db.Scopes(Paginate(r)).Table("articles").Find(&articles)
	repo.db.Table("articles").Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, article := range articles {
		articlesDTO = append(articlesDTO, article.ToArticleDTO())
	}

	return articlesDTO, totalElements, nil
}

func (repo *Repository) FindAllArticlesFromRestaurant(r *http.Request) ([]models.ArticleDTO, int64, error) {
	var articlesDTO []models.ArticleDTO
	var articles []*models.Article
	var totalElements int64

	restaurantName := r.URL.Query().Get("restaurantName")

	result := repo.db.Scopes(Paginate(r)).Table("articles").
		Where("deleted_at IS NULL and restaurant_name = ?", restaurantName).
		Find(&articles)

	repo.db.Table("articles").
		Where("deleted_at IS NULL and restaurant_name = ?", restaurantName).
		Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, article := range articles {
		articlesDTO = append(articlesDTO, article.ToArticleDTO())
	}

	return articlesDTO, totalElements, nil
}

func (repo *Repository) SearchArticles(r *http.Request) ([]models.ArticleDTO, int64, error) {
	var articlesDTO []models.ArticleDTO
	var articles []*models.Article
	var totalElements int64

	restaurantName := r.URL.Query().Get("restaurantName")
	searchField := r.URL.Query().Get("searchField")
	articleType := r.URL.Query().Get("articleType")
	priceFrom := r.URL.Query().Get("priceFrom")
	priceTo := r.URL.Query().Get("priceTo")

	result := repo.db.Scopes(Paginate(r)).Table("articles").
		Where("(deleted_at IS NULL and restaurant_name = ?) and "+
			"('' = ? or LOWER(article_name) LIKE ?) and "+
			"('' = ? or article_type = ?) and "+
			"(price >= ? and price <= ?)",
			restaurantName,
			searchField, concat(searchField), articleType, articleType,
			priceFrom, priceTo).
		Find(&articles)

	repo.db.Table("articles").
		Where("(deleted_at IS NULL and restaurant_name = ?) and "+
			"('' = ? or LOWER(article_name) LIKE ?) and "+
			"('' = ? or article_type = ?) and "+
			"(price >= ? and price <= ?)",
			restaurantName,
			searchField, concat(searchField), articleType, articleType,
			priceFrom, priceTo).
		Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, article := range articles {
		articlesDTO = append(articlesDTO, article.ToArticleDTO())
	}

	return articlesDTO, totalElements, nil
}

func (repo *Repository) FindArticleById(id uint) (*models.ArticleDTO, error) {
	var article models.Article
	result := repo.db.Table("articles").Where("id = ?", id).First(&article)

	if result.Error != nil {
		return nil, errors.New("article cannot be found")
	}

	var retValue models.ArticleDTO = article.ToArticleDTO()
	return &retValue, nil
}

func (repo *Repository) CreateArticle(articleDTO *models.ArticleDTO) (*models.ArticleDTO, error) {
	var article models.Article = articleDTO.ToArticle()
	result := repo.db.Table("articles").Create(&article)

	if result.Error != nil {
		return nil, errors.New("error while creating article")
	}

	var retValue models.ArticleDTO = article.ToArticleDTO()
	return &retValue, nil
}

func (repo *Repository) UpdateArticle(articleDTO *models.ArticleDTO) (*models.ArticleDTO, error) {
	var article models.Article
	result := repo.db.Table("articles").Where("id = ?", articleDTO.Id).First(&article)

	if result.Error != nil {
		return nil, errors.New("article cannot be found")
	}

	if articleDTO.Changed {
		article.Image = articleDTO.ImagePath
	}

	article.ArticleName = articleDTO.ArticleName
	article.ArticleType = articleDTO.ArticleType
	article.Price = articleDTO.Price
	article.Description = articleDTO.Description
	article.RestaurantName = articleDTO.RestaurantName

	result2 := repo.db.Table("articles").Save(&article)

	if result2.Error != nil {
		return nil, errors.New("error while updating article")
	}

	var retValue models.ArticleDTO = article.ToArticleDTO()
	return &retValue, nil
}

func (repo *Repository) DeleteArticle(id uint) (*models.ArticleDTO, error) {
	var article models.Article
	result := repo.db.Table("articles").Where("id = ?", id).Clauses(clause.Returning{}).Delete(&article)

	if result.Error != nil {
		return nil, errors.New("error while deleting article")
	}

	var retValue models.ArticleDTO = article.ToArticleDTO()
	return &retValue, nil
}
