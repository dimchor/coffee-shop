package controller

import (
	"coffee_shop_backend/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	// temporary, GORM will be used in the (near) future
	coffeebeans []product.CoffeeBean
	cups        []product.Cup
}

func NewProductController() *ProductController {
	newProductController := &ProductController{
		coffeebeans: [](product.CoffeeBean){
			product.CoffeeBean{
				product.ProductBase{1, "Colombia Las Flores", 9.8},
				"Arabica Red Bourbon",
				"Acevedo, Huila",
				200},
			product.CoffeeBean{
				product.ProductBase{2, "Brazil Alta Mogiana", 6.8},
				"Arabica Catuai, Mundo novo, Acaia",
				"Alta Mogiana, São Paulo",
				200},
		},
		cups: [](product.Cup){
			product.Cup{
				product.ProductBase{1, "Small cup", 1.5},
				product.SMALL},
			product.Cup{
				product.ProductBase{2, "Medium cup", 2.5},
				product.MEDIUM},
			product.Cup{
				product.ProductBase{3, "Large cup", 3.5},
				product.LARGE},
		},
	}
	return newProductController
}

func (p *ProductController) GetCoffeeBeans(c *gin.Context) {
	c.JSON(http.StatusOK, p.coffeebeans)
}

func (p *ProductController) GetCups(c *gin.Context) {
	c.JSON(http.StatusOK, p.cups)
}

func (p *ProductController) GetCoffeeBeansById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		// TODO: return appropriate code for invalid input
		return
	}

	// TODO: To be replaced by gorm call
	for _, item := range p.coffeebeans {
		if item.Id == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "coffee beans not found"})
}

func (p *ProductController) GetCupsById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		// TODO: return appropriate code for invalid input
		return
	}

	// TODO: To be replaced by gorm call
	for _, item := range p.cups {
		if item.Id == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "cup not found"})
}

func (p *ProductController) PostCoffeeBeans(c *gin.Context) {
	var coffeeBean product.CoffeeBean

	if err := c.BindJSON(&coffeeBean); err != nil {
		// TODO: find a proper return code for failure to bind JSON
		return
	}

	// TODO: To be replaced by gorm call
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

func (p *ProductController) PostCup(c *gin.Context) {
	var cup product.Cup

	if err := c.BindJSON(&cup); err != nil {
		// TODO: find a proper return code for failure to bind JSON
		return
	}

	// TODO: To be replaced by gorm call
	for _, item := range p.cups {
		if item.Id == cup.Id {
			// TODO: find a proper return code for same id
			return
		}
	}

	// perform other checks here, if needed

	p.cups = append(p.cups, cup)
	c.JSON(http.StatusCreated, cup)
}
