package models

import "gorm.io/gorm"

func (review *Review) ToReviewDTO() ReviewDTO {
	return ReviewDTO{
		Id:                   review.ID,
		Comment:              review.Comment,
		Rating:               review.Rating,
		InappropriateContent: review.InappropriateContent,
		DateTime:             review.DateTime,
		IdRestaurant:         review.IdRestaurant,
		IdOrder:              review.IdOrder,
		IdUser:               review.IdUser,
		EmailUser:            review.EmailUser,
	}
}

func (reviewDTO *ReviewDTO) ToReview() Review {
	return Review{
		Model:                gorm.Model{},
		Comment:              reviewDTO.Comment,
		Rating:               reviewDTO.Rating,
		InappropriateContent: reviewDTO.InappropriateContent,
		DateTime:             reviewDTO.DateTime,
		IdRestaurant:         reviewDTO.IdRestaurant,
		IdOrder:              reviewDTO.IdOrder,
		IdUser:               reviewDTO.IdUser,
		EmailUser:            reviewDTO.EmailUser,
	}
}
