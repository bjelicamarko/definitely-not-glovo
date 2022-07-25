package models

type ArticleDTO struct {
	Id             uint    `json:"Id"`
	ArticleName    string  `json:"ArticleName"`
	ArticleType    string  `json:"ArticleType"`
	Price          float32 `json:"Price"`
	Description    string  `json:"Description"`
	RestaurantName string  `json:"RestaurantName"`
	Image          string  `json:"Image"`
	ImagePath      string  `json:"ImagePath"`
	Changed        bool    `json:"Changed"`
}

type ArticleDTOMessage struct {
	ArticleDTO ArticleDTO `json:"ArticleDTO"`
	Message    string     `json:"Message"`
}

type ArticlesPageable struct {
	Elements      []ArticleDTO `json:"Elements"`
	TotalElements int64        `json:"TotalElements"`
}
