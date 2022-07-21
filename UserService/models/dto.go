package models

type UserDTO struct {
	Id        uint   `json:"Id"`
	Email     string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Contact   string `json:"Contact"`
	Role      string `json:"Role"`
	Banned    bool   `json:"Banned"`
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
