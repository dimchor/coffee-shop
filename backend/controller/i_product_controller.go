package controller

import "github.com/gin-gonic/gin"

type IProductController interface {
	GetCoffeeBeans(*gin.Context)
	GetCoffeeBeansById(c *gin.Context)
	PostCoffeeBeans(c *gin.Context)
	GetCups(c *gin.Context)
	GetCupsById(c *gin.Context)
	PostCup(c *gin.Context)
	GetCoffeeDrink(c *gin.Context)
	GetCoffeeDrinkById(c *gin.Context)
	PostCoffeeDrink(c *gin.Context)
}
