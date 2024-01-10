package types

import "strconv"

type Product struct {
	Base
	Name        string
	Price       float64
	Description string
}

type ProductDto struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

func (p *Product) ToProductDto() *ProductDto {
	return &ProductDto{strconv.FormatUint(p.ID, 10), p.Name, p.Price,
		p.Description}
}

func (p *ProductDto) ToProduct() (*Product, error) {
	id, err := strconv.ParseUint(p.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	var base Base
	base.ID = id
	return &Product{base, p.Name, p.Price, p.Description}, nil
}
