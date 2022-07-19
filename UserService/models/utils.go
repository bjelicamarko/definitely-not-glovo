package models

import (
	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Message string `json:"message"`
}

type Credentials struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Claims struct {
	Email string `json:"email"`
	Role  Role   `json:"role"`
	Id    uint   `json:"Id"`
	jwt.StandardClaims
}

type LoginResponse struct {
	Token string `json:"Token"`
}
