package service

import (
	"coffee_shop_backend/product"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) (*ProductService, error) {
	err := db.AutoMigrate(&product.CoffeeBean{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&product.CoffeeDrink{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&product.Cup{})
	if err != nil {
		return nil, err
	}

	return &ProductService{db}, nil
}

func (p *ProductService) GetCoffeeBeans() ([]product.CoffeeBean, error) {
	var beans []product.CoffeeBean
	result := p.db.Find(&beans)
	return beans, result.Error
}

func (p *ProductService) GetCoffeeBeanById(id uint64) (
	*product.CoffeeBean, error) {
	var bean product.CoffeeBean
	result := p.db.First(&bean, id)
	return &bean, result.Error
}

func (p *ProductService) PostCoffeeBean(bean *product.CoffeeBean) error {
	result := p.db.Create(bean)
	return result.Error
}

func (p *ProductService) GetCups() ([]product.Cup, error) {
	var cups []product.Cup
	result := p.db.Find(&cups)
	return cups, result.Error
}

func (p *ProductService) GetCupById(id uint64) (*product.Cup, error) {
	var cup product.Cup
	result := p.db.First(&cup, id)
	return &cup, result.Error
}

func (p *ProductService) PostCup(cup *product.Cup) error {
	result := p.db.Create(cup)
	return result.Error
}

func (p *ProductService) GetCoffeeDrinks() ([]product.CoffeeDrink, error) {
	var drinks []product.CoffeeDrink
	result := p.db.Find(&drinks)
	return drinks, result.Error
}

func (p *ProductService) GetCoffeeDrinkById(id uint64) (
	*product.CoffeeDrink, error) {
	var drink product.CoffeeDrink
	result := p.db.First(&drink, id)
	return &drink, result.Error
}

func (p *ProductService) PostCoffeeDrink(drink *product.CoffeeDrink) error {
	result := p.db.Create(drink)
	return result.Error
}
