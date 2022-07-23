package models

import "UserService/utils"

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
	}
}
