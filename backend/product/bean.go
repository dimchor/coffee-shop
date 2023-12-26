package product

type CoffeeBean struct {
	ProductBase `json:"product"`
	Species     string  `json:"species"`
	Area        string  `json:"area"`
	Weight      float64 `json:"weight"`
}
