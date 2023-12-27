package service

import (
	"coffee_shop_backend/product"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	db.AutoMigrate(&product.CoffeeBean{})
	db.AutoMigrate(&product.CoffeeDrink{})
	db.AutoMigrate(&product.Cup{})
	return &ProductService{db}
}

func (*ProductService) GetCoffeeBeans() {

}
func (*ProductService) GetCoffeeBeansById() {

}
func (*ProductService) PostCoffeeBeans() {

}
func (*ProductService) GetCups() {

}
func (*ProductService) GetCupsById() {

}
func (*ProductService) PostCup() {

}
func (*ProductService) GetCoffeeDrink() {

}
func (*ProductService) GetCoffeeDrinkById() {

}
func (*ProductService) PostCoffeeDrink() {

}
