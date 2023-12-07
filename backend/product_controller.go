package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	// temporary, GORM will be used in the (near) future
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

func (p *ProductController) GetCoffeeBeansById(c *gin.Context) {
	id := c.Param("id")

	for _, item := range p.coffeebeans {
		if item.Id == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "coffee beans not found"})
}

func (p *ProductController) PostCoffeeBeans(c *gin.Context) {
	var coffeeBean CoffeeBean

	if err := c.BindJSON(&coffeeBean); err != nil {
		// TODO: find a proper return code for failure to bind JSON
		return
	}

	for _, item := range p.coffeebeans {
		if item.Id == coffeeBean.Id {
			// TODO: find a proper return code for same id
			return
		}
	}

	// perform other checks here, if needed

	p.coffeebeans = append(p.coffeebeans, coffeeBean)
	c.JSON(http.StatusCreated, coffeeBean)
}
