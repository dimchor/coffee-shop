package service

type IProductService interface {
	GetCoffeeBeans()
	GetCoffeeBeansById()
	PostCoffeeBeans()
	GetCups()
	GetCupsById()
	PostCup()
	GetCoffeeDrink()
	GetCoffeeDrinkById()
	PostCoffeeDrink()
}
