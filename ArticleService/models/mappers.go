package models

import (
	"ArticleService/utils"

	"gorm.io/gorm"
)

func (article *Article) ToArticleDTO() ArticleDTO {
	return ArticleDTO{
		Id:             article.ID,
		ArticleName:    article.ArticleName,
		ArticleType:    article.ArticleType,
		Price:          article.Price,
		Description:    article.Description,
		RestaurantName: article.RestaurantName,
		Image:          utils.GetB64Image(article.Image),
		ImagePath:      article.Image,
		Changed:        false,
	}
}

func (articleDTO *ArticleDTO) ToArticle() Article {
	return Article{
		Model:          gorm.Model{},
		ArticleName:    articleDTO.ArticleName,
		ArticleType:    articleDTO.ArticleType,
		Price:          articleDTO.Price,
		Description:    articleDTO.Description,
		RestaurantName: articleDTO.RestaurantName,
		Image:          articleDTO.ImagePath,
	}
}
