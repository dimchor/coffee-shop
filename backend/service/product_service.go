package service

import (
	"coffee_shop_backend/types"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) (*ProductService, error) {
	err := db.AutoMigrate(&types.Product{})
	if err != nil {
		return nil, err
	}

	return &ProductService{db}, nil
}

func (p *ProductService) GetProducts() ([]types.Product, error) {
	var products []types.Product
	result := p.db.Find(&products)
	return products, result.Error
}

func (p *ProductService) GetProductById(id uint64) (*types.Product, error) {
	var product types.Product
	result := p.db.First(&product, id)
	return &product, result.Error
}

func (p *ProductService) PostProduct(product *types.Product) error {
	result := p.db.Create(product)
	return result.Error
}
