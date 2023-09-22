package services

import (
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/entities"
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/interfaces"
)

type UserService struct {
}

func InitUserService() interfaces.IUser {
	return &UserService{}
}
func (u *UserService) GetLogin(user *entities.User) string {
	return "user logged in"

}

func (u *UserService) GetLogout(userId int) string {
	return "user logged out"

}

func (u *UserService) GetRegister(user *entities.User) string {
	return "user registered"

}

func (u *UserService) GetProfile(userId int) *entities.User {
	return &entities.User{
		UserId:   1,
		Name:     "ARchi",
		Email:    "a@gmail.com",
		Password: "abc"}

}
