package repository

import "Challenge_11/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() *entity.Product
}
