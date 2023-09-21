package services

import (
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/entities"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/interfaces"
)

type UserService struct {
}

func InitUserService() interfaces.IUser {
	return &UserService{}
}

func (u *UserService) Login(user *entities.User) string {
	return "User Logged in"

}

func (u *UserService) Register(user *entities.User) string {
	return "user registered"
}

func (u *UserService) GetProfile(userId int) *entities.User {
	return &entities.User{
		Name:     "Archi",
		UserId:   1,
		Password: "archi",
		Email:    "archi@gmail.com",
		Role:     "Lead"}
}

func (u *UserService) SearchUser(query string) {

}

func (u *UserService) Logout(userid int) string {
	return "user logged out"

}
