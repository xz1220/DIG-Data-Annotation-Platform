package dto

import "labelproject-back/model"

//UserDto is user information which is should be returned
type UserDto struct {
	Username    string `json:"username"`
	UserId      int64  `json:"userId"`
	Authorities string `json:"authorities"`
}

// ToUserDto is a kind of translation
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Username:    user.Username,
		UserId:      user.UserID,
		Authorities: user.Authorities,
	}
}
