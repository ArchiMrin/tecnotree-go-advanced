package interfaces

import "github.com/ArchiMrin/tecnotree-go-advanced/ekart_1/entities"

type IProduct interface {
	Insert(product *entities.Product) (string, error)
}
