package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
  	"gorm.io/driver/mysql"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// for testing purposes
type Help struct {
	gorm.Model
	Help_topic_id uint
	Name string
	Help_category_id uint
	Description string
	Example string
	Url string
}

func main() {
	//cup := Cup{Product{"Standard cup", "c0001", 1.0}, MEDIUM}

	//fmt.Println(cup)

	password := os.Getenv("DB_ROOT_PASSWORD")
	database_name := "mysql"
	dsn := "root:" + password + "@tcp(db:3306)/" +
		database_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Coudln't connect to database!")
		return
	}

	var help Help
	database.First(&help)
	fmt.Println(help.Description)

	productController := NewProductController()

	r := gin.Default()
	r.GET("/ping", ping)

	r.GET("/getcoffeebeans", productController.GetCoffeeBeans)
	r.GET("/getcoffeebean/:id", productController.GetCoffeeBeansById)
	r.POST("/postcoffeebeans", productController.PostCoffeeBeans)

	r.GET("/getcups", productController.GetCups)
	r.GET("/getcup/:id", productController.GetCupsById)
	r.POST("/postcup", productController.PostCup)
	r.Run()
}
