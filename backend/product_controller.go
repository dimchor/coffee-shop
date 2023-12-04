package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	coffeebeans []CoffeeBean
}

func NewProductController() *ProductController {
	newProductController := &ProductController{
		coffeebeans: [](CoffeeBean){
			CoffeeBean{
				Product{"Colombia Las Flores", "B1", 9.8},
				"Arabica Red Bourbon",
				"Acevedo, Huila",
				200},
			CoffeeBean{
				Product{"Brazil Alta Mogiana", "B2", 6.8},
				"Arabica Catuai, Mundo novo, Acaia",
				"Alta Mogiana, São Paulo",
				200},
		},
	}
	return newProductController
}

func (p *ProductController) GetCoffeeBeans(c *gin.Context) {
	c.JSON(http.StatusOK, p.coffeebeans)
}
