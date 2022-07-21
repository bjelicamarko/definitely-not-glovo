package handlers

import (
	"UserService/models"
	"UserService/repository"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type UsersHandler struct {
	repository *repository.Repository
}

func NewUsersHandler(repository *repository.Repository) *UsersHandler {
	return &UsersHandler{repository}
}

var jwtKey = []byte("z7031Q8Qy9zVO-T2o7lsFIZSrd05hH0PaeaWIBvLh9s")

func (uh *UsersHandler) Login(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var credentials models.Credentials
	json.NewDecoder(req.Body).Decode(&credentials)

	user, err := uh.repository.CheckCredentials(credentials.Email, credentials.Password)

	if err != nil {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: err.Error()})
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := models.Claims{Email: user.Email, Role: user.Role, Id: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(resWriter).Encode(models.LoginResponse{Token: tokenString})
}

func Authorize(r *http.Request) (*jwt.Token, error) {
	cookie := r.Header.Values("Authorization")
	tokenString := strings.Split(cookie[0], " ")[1]

	claims := models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	return token, err
}

func (uh *UsersHandler) AuthorizeAdmin(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.ADMIN {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) AuthorizeAppUser(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.APPUSER {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) AuthorizeEmployee(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.EMPLOYEE {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) AuthorizeDeliverer(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.DELIVERER {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) Register(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var newUserDTO models.NewUserDTO
	json.NewDecoder(req.Body).Decode(&newUserDTO)

	_, err := uh.repository.CreateUser(&newUserDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "registration failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "registration succeeded"})
}

func (uh *UsersHandler) GetUsers(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	users, totalElements, _ := uh.repository.FindAll(req)

	json.NewEncoder(resWriter).Encode(models.UsersPageable{Elements: users, TotalElements: totalElements})
}

func (uh *UsersHandler) UpdateUser(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	var updatedUser models.UserDTO
	json.NewDecoder(req.Body).Decode(&updatedUser)

	_, err := uh.repository.UpdateUser(&updatedUser)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: err.Error()})
	} else {
		json.NewEncoder(resWriter).Encode(models.Response{Message: "User successfully updated"})
	}
}

func (uh *UsersHandler) DeleteUser(resWriter http.ResponseWriter, req *http.Request) {
	AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	err := uh.repository.DeleteUser(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "error while deleting user"})
	} else {
		json.NewEncoder(resWriter).Encode(models.Response{Message: "user successfully deleted"})
	}
}
