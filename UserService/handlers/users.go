package handlers

import (
	"UserService/models"
	"UserService/repository"
	"UserService/utils"
	"encoding/json"
	"net/http"
	"os"
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
	utils.AdjustResponseHeaderJson(&resWriter)

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
	utils.AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.ADMIN {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) AuthorizeAppUser(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.APPUSER {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) AuthorizeEmployee(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.EMPLOYEE {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) AuthorizeDeliverer(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	token, err := Authorize(req)

	if err != nil || !token.Valid || token.Claims.(*models.Claims).Role != models.DELIVERER {
		resWriter.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "authorization succeeded"})
}

func (uh *UsersHandler) Register(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	var newUserDTO models.UserDTO
	json.NewDecoder(req.Body).Decode(&newUserDTO)

	err := uh.repository.Register(&newUserDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.Response{Message: "registration failed"})
		return
	}

	json.NewEncoder(resWriter).Encode(models.Response{Message: "registration succeeded"})
}

func (uh *UsersHandler) FindAllUsers(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	usersDTO, totalElements, _ := uh.repository.FindAllUsers(req)

	json.NewEncoder(resWriter).Encode(models.UsersPageable{Elements: usersDTO, TotalElements: totalElements})
}

func (uh *UsersHandler) SearchUsers(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	usersDTO, totalElements, _ := uh.repository.SearchUsers(req)

	json.NewEncoder(resWriter).Encode(models.UsersPageable{Elements: usersDTO, TotalElements: totalElements})
}

func (uh *UsersHandler) FindUserById(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := uh.repository.FindUserById(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: "user successfully found", UserDTO: *userDTO})
}

func (uh *UsersHandler) CreateUser(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	var newUserDTO models.UserDTO
	json.NewDecoder(req.Body).Decode(&newUserDTO)

	_ = os.Remove(newUserDTO.ImagePath)
	utils.ToImage(newUserDTO.Image, newUserDTO.ImagePath)

	userDTO, err := uh.repository.CreateUser(&newUserDTO)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: "user successfully created", UserDTO: *userDTO})
}

func (uh *UsersHandler) UpdateUser(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	var updatedUser models.UserDTO
	json.NewDecoder(req.Body).Decode(&updatedUser)

	if updatedUser.Changed {
		_ = os.Remove(updatedUser.ImagePath)
		utils.ToImage(updatedUser.Image, updatedUser.ImagePath)
	}

	userDTO, err := uh.repository.UpdateUser(&updatedUser)

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: "user successfully updated", UserDTO: *userDTO})
}

func (uh *UsersHandler) DeleteUser(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := uh.repository.DeleteUser(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: "user successfully deleted", UserDTO: *userDTO})
}

func (uh *UsersHandler) BanUser(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := uh.repository.BanUser(uint(idInt))
	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: "user successfully banned", UserDTO: *userDTO})
}

func (uh *UsersHandler) UnbanUser(resWriter http.ResponseWriter, req *http.Request) {
	utils.AdjustResponseHeaderJson(&resWriter)

	params := mux.Vars(req)
	idStr := params["id"]
	idInt, _ := strconv.ParseInt(idStr, 10, 64)

	userDTO, err := uh.repository.UnbanUser(uint(idInt))

	if err != nil {
		resWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: err.Error()})
		return
	}

	json.NewEncoder(resWriter).Encode(models.UserDTOMessage{Message: "user successfully unbanned", UserDTO: *userDTO})
}
