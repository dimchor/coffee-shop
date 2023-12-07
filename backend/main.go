package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func main() {
	//cup := Cup{Product{"Standard cup", "c0001", 1.0}, MEDIUM}

	//fmt.Println(cup)

	fmt.Println("hello, world")

	productController := NewProductController()

	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/getcoffeebeans", productController.GetCoffeeBeans)
	r.GET("/getcoffeebeans/:id", productController.GetCoffeeBeansById)
	r.POST("/postcoffeebeans", productController.PostCoffeeBeans)
	r.Run()
}
