package interfaces

import (
	"github.com/ArchiMrin/tecnotree-go-advanced/ekart/entities"
)

type IUser interface {
	Register(user *entities.User) string
	Login(user *entities.User) string
	GetProfile(userId int) *entities.User
	SearchUser(searchQUery string)
	Logout(userId int) string //return type string
}
