package repository

import (
	"ReviewService/models"
	"errors"
	"net/http"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
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

func (repo *Repository) GetReviewsOfRestaurant(r *http.Request) ([]models.ReviewDTO, int64, error) {
	var reviewsDTO []models.ReviewDTO
	var reviews []*models.Review
	var totalElements int64

	restaurantId := r.URL.Query().Get("restaurantId")

	result := repo.db.Scopes(Paginate(r)).Table("reviews").
		Where("deleted_at is null and id_restaurant = ?",
			restaurantId).
		Order("id desc").
		Find(&reviews)

	repo.db.Table("reviews").
		Where("deleted_at is null and id_restaurant = ?",
			restaurantId).
		Order("id desc").
		Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, review := range reviews {
		reviewsDTO = append(reviewsDTO, review.ToReviewDTO())
	}

	return reviewsDTO, totalElements, nil
}

func (repo *Repository) SearchReviews(r *http.Request) ([]models.ReviewDTO, int64, error) {
	var reviewsDTO []models.ReviewDTO
	var reviews []*models.Review
	var totalElements int64

	restaurantId := r.URL.Query().Get("restaurantId")
	userId := r.URL.Query().Get("userId")
	inappropriate := r.URL.Query().Get("inappropriate")

	if restaurantId == "" {
		restaurantId = "0"
	}

	if userId == "" {
		userId = "0"
	}

	if inappropriate == "" {
		inappropriate = "false"
	}

	result := repo.db.Scopes(Paginate(r)).Table("reviews").
		Where("deleted_at is null and "+
			"('0' = ? or id_restaurant = ?) and "+
			"('0' = ? or id_user = ?) and "+
			"('' = ? or inappropriate_content = ?)",
			restaurantId, restaurantId, userId, userId, inappropriate, inappropriate).
		Order("id desc").
		Find(&reviews)

	repo.db.Table("reviews").
		Where("deleted_at is null and "+
			"('0' = ? or id_restaurant = ?) and "+
			"('0' = ? or id_user = ?) and "+
			"('' = ? or inappropriate_content = ?)",
			restaurantId, restaurantId, userId, userId, inappropriate, inappropriate).
		Order("id desc").
		Count(&totalElements)

	if result.Error != nil {
		return nil, totalElements, result.Error
	}

	for _, review := range reviews {
		reviewsDTO = append(reviewsDTO, review.ToReviewDTO())
	}

	return reviewsDTO, totalElements, nil
}

func (repo *Repository) FindReviewByOrder(id uint) (*models.ReviewDTO, error) {
	var review models.Review
	result := repo.db.Table("reviews").Where("id_order = ?", id).First(&review)

	if result.Error != nil {
		return nil, errors.New("review cannot be found")
	}

	var reviewDTO models.ReviewDTO = review.ToReviewDTO()
	return &reviewDTO, nil
}

func (repo *Repository) CreateReview(reviewDTO *models.ReviewDTO) (*models.ReviewDTO, error) {
	var review models.Review = reviewDTO.ToReview()
	result := repo.db.Table("reviews").Create(&review)

	if result.Error != nil {
		return nil, errors.New("error while creating review")
	}

	var retValue models.ReviewDTO = review.ToReviewDTO()
	return &retValue, nil
}

func (repo *Repository) ReportReview(reviewDTO *models.ReviewDTO) (*models.ReviewDTO, error) {
	var review models.Review
	result := repo.db.Table("reviews").Where("id = ?", reviewDTO.Id).First(&review)

	if result.Error != nil {
		return nil, errors.New("review cannot be found")
	}

	review.InappropriateContent = true

	result2 := repo.db.Table("reviews").Save(&review)

	if result2.Error != nil {
		return nil, errors.New("error while reporting review")
	}

	var retValue models.ReviewDTO = review.ToReviewDTO()
	return &retValue, nil
}

func (repo *Repository) DeleteReview(id uint) (*models.ReviewDTO, error) {
	var review models.Review
	result := repo.db.Where("id = ?", id).Clauses(clause.Returning{}).Delete(&review)

	if result.Error != nil {
		return nil, errors.New("error while deleting review")
	}

	var retValue models.ReviewDTO = review.ToReviewDTO()
	return &retValue, nil
}
