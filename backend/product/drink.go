package product

type CoffeeDrink struct {
	ProductBase `json:"product"`
	CoffeeBean  `json:"bean"`
	Iced        bool `json:"iced"`
}
