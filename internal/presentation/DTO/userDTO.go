package DTO

type UserDTO struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessLevel int    `json:"access_level"`
}

//func ToUserDTO(user *entities.User) UserDTO {
//	return UserDTO{
//		Username: user.Username,
//		Password:
//	}
//}
