package models

import (
	"UserService/utils"

	"gorm.io/gorm"
)

func (user *User) ToUserDTO() UserDTO {
	return UserDTO{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Contact:   user.Contact,
		Role:      string(user.Role),
		Banned:    user.Banned,
		Image:     utils.GetB64Image(user.Image),
		ImagePath: user.Image,
		Changed:   false,
	}
}

func (userDTO *UserDTO) ToUser() User {
	return User{
		Model:     gorm.Model{},
		Email:     userDTO.Email,
		Password:  userDTO.Password,
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Contact:   userDTO.Contact,
		Role:      Role(userDTO.Role),
		Banned:    userDTO.Banned,
		Image:     userDTO.ImagePath,
	}
}
