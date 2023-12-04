package main

type CoffeeBean struct {
	Product `json:"product"`
	Species string  `json:"species"`
	Area    string  `json:"area"`
	Weight  float64 `json:"weight"`
}
