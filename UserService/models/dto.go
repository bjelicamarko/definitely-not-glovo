package models

type UserDTO struct {
	Id             uint   `json:"Id"`
	Email          string `json:"Email"`
	Password       string `json:"Password"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	Contact        string `json:"Contact"`
	Role           string `json:"Role"`
	Banned         bool   `json:"Banned"`
	Image          string `json:"Image"`
	ImagePath      string `json:"ImagePath"`
	Changed        bool   `json:"Changed"`
	RestaurantName string `json:"RestaurantName"`
}

type UserDTOMessage struct {
	UserDTO UserDTO `json:"UserDTO"`
	Message string  `json:"Message"`
}

type UsersPageable struct {
	Elements      []UserDTO `json:"Elements"`
	TotalElements int64     `json:"TotalElements"`
}
