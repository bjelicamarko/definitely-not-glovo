package models

type RestaurantDTO struct {
	Id             uint    `json:"Id"`
	RestaurantName string  `json:"RestaurantName"`
	ContactPhone   string  `json:"ContactPhone"`
	Image          string  `json:"Image"`
	ImagePath      string  `json:"ImagePath"`
	Country        string  `json:"Country"`
	City           string  `json:"City"`
	Street         string  `json:"Street"`
	StreetNumber   string  `json:"StreetNumber"`
	Ptt            uint    `json:"Ptt"`
	DisplayName    string  `json:"DisplayName"`
	Longitude      float32 `json:"Longitude"`
	Latitude       float32 `json:"Latitude"`
	Changed        bool    `json:"Changed"`
}

type RestaurantDTOMessage struct {
	RestaurantDTO RestaurantDTO `json:"RestaurantDTO"`
	Message       string        `json:"Message"`
}

type RestaurantsPageable struct {
	Elements      []RestaurantDTO `json:"Elements"`
	TotalElements int64           `json:"TotalElements"`
}
