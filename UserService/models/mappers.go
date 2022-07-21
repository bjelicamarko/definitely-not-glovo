package models

func (user *User) ToUserDTO() UserDTO {
	return UserDTO{
		Id:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Contact:   user.Contact,
		Role:      string(user.Role),
		Banned:    user.Banned,
	}
}
