package models

type UserDTO struct {
	Id        uint   `json:"Id"`
	Email     string `json:"Email"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Contact   string `json:"Contact"`
}

type NewUserDTO struct {
	Email     string `json:"Email"`
	Password  string `json:"Password"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Contact   string `json:"Contact"`
}
