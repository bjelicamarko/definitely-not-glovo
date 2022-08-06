package models

type ReviewDTO struct {
	Id                   uint   `json:"Id"`
	Comment              string `json:"Comment"`
	Rating               int    `json:"Rating"`
	InappropriateContent bool   `json:"InappropriateContent"`
	DateTime             string `json:"DateTime"`
	IdRestaurant         uint   `json:"IdRestaurant"`
	IdOrder              uint   `json:"IdOrder"`
	IdUser               uint   `json:"IdUser"`
	EmailUser            string `json:"EmailUser"`
}

type ReviewDTOMessage struct {
	ReviewDTO ReviewDTO `json:"ReviewDTO"`
	Message   string    `json:"Message"`
}

type ReviewsPageable struct {
	Elements      []ReviewDTO `json:"Elements"`
	TotalElements int64       `json:"TotalElements"`
}
