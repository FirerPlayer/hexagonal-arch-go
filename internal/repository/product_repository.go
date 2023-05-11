package repository

import "github.com/firerplayer/hexagonal-arch-go/internal/entity"

type ProductRepository interface {
	Create(product *entity.Product) error
	FindAll() ([]*entity.Product, error)
}
