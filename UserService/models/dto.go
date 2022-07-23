package models

type UserDTO struct {
	Id        uint   `json:"Id"`
	Email     string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Contact   string `json:"Contact"`
	Role      string `json:"Role"`
	Banned    bool   `json:"Banned"`
	Image     string `json:"Image"`
}

type NewUserDTO struct {
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Contact   string `json:"Contact"`
}

type UsersPageable struct {
	Elements      []UserDTO `json:"Elements"`
	TotalElements int64     `json:"TotalElements"`
}

type ImageMessage struct {
	Image string `json:"Image"`
	Path  string `json:"Path"`
	Id    uint   `json:"Id"`
}

type UserDTOMessage struct {
	UserDTO UserDTO `json:"UserDTO"`
	Message string  `json:"Message"`
}
