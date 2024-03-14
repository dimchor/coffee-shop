package service

import "coffee_shop_backend/types"

type IProductService interface {
	GetProducts() ([]types.Product, error)
	GetProductById(uint64) (*types.Product, error)
	PostProduct(*types.Product) error
	PostNewUser(*types.UserCreateDto) error
	PostLoginUser(*types.UserLoginDto) (string, error)
}
