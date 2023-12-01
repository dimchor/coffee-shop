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
	cup := Cup{Product{"Standard cup", "c0001", 1.0}, MEDIUM}

	bean := CoffeeBean{Product{"Colombia Las Flores", "1234", 9.8}, "Arabica Red Bourbon", "Acevedo, Huila, Brazil"}

	fmt.Println(bean)
	fmt.Println(cup)

	fmt.Println("hello, world")

	r := gin.Default()
	r.GET("/ping", ping)
	r.Run()
}
