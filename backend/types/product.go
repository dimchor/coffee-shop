package types

import "strconv"

type Product struct {
	Base
	Name        string
	Price       float64
	Description string
}

type ProductGetDto struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

type ProductPostDto struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Token       string  `json:"token"`
}

func (p *Product) ToProductGetDto() *ProductGetDto {
	return &ProductGetDto{strconv.FormatUint(p.ID, 10), p.Name, p.Price,
		p.Description}
}

func (p *ProductPostDto) ToProduct() *Product {
	return &Product{Base{}, p.Name, p.Price, p.Description}
}
