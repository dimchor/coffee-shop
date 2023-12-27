package controller

import "github.com/gin-gonic/gin"

type IProductController interface {
	GetCoffeeBeans(*gin.Context)
	GetCoffeeBeanById(c *gin.Context)
	PostCoffeeBean(c *gin.Context)
	GetCups(c *gin.Context)
	GetCupById(c *gin.Context)
	PostCup(c *gin.Context)
	GetCoffeeDrinks(c *gin.Context)
	GetCoffeeDrinkById(c *gin.Context)
	PostCoffeeDrink(c *gin.Context)
}
