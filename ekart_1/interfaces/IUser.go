package interfaces

import "github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/entities"

type IUser interface {
	GetRegister(user *entities.User) string
	GetLogin(user *entities.User) string
	GetProfile(userId int) *entities.User
	GetLogout(userId int) string
}
