package service

import "coffee_shop_backend/product"

type IProductService interface {
	GetCoffeeBeans() ([]product.CoffeeBean, error)
	GetCoffeeBeanById(uint64) (*product.CoffeeBean, error)
	PostCoffeeBean(*product.CoffeeBean) error
	GetCups() ([]product.Cup, error)
	GetCupById(uint64) (*product.Cup, error)
	PostCup(*product.Cup) error
	GetCoffeeDrinks() ([]product.CoffeeDrink, error)
	GetCoffeeDrinkById(uint64) (*product.CoffeeDrink, error)
	PostCoffeeDrink(*product.CoffeeDrink) error
}
