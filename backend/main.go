package main

import (
	"coffee_shop_backend/controller"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

// for testing purposes
type Tabler interface {
	TableName() string
}

type Help struct {
	HelpTopicId    uint   `gorm:"primaryKey"`
	Name           string `gorm:"size:64"`
	HelpCategoryId uint
	Description    string
	Example        string
	Url            string
}

func (Help) TableName() string {
	return "help_topic"
}

// ...

type WaitOptions struct {
	Interval    uint
	StartPeriod uint
	Retries     uint
}

func try_opening(dsn string, config *gorm.Config,
	options WaitOptions) (db *gorm.DB, err error) {

	fmt.Printf("Waiting for %d seconds...\n", options.StartPeriod)
	time.Sleep(time.Duration(options.StartPeriod) * time.Second)
	for ; options.Retries > 0; options.Retries-- {
		db, err = gorm.Open(mysql.Open(dsn), config)
		if err == nil {
			fmt.Printf("Connected to database successfully!\n")
			return
		}
		fmt.Printf("Failed to connect to database. Retring in %d seconds...\n",
			options.Interval)
		time.Sleep(time.Duration(options.Interval) * time.Second)
	}
	fmt.Printf("Failed to connect to database. Aborting....\n")
	return
}

func main() {
	//cup := Cup{Product{"Standard cup", "c0001", 1.0}, MEDIUM}

	//fmt.Println(cup)

	password := os.Getenv("DB_ROOT_PASSWORD")
	const database_name = "mysql"
	dsn := "root:" + password + "@tcp(db:3306)/" +
		database_name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := try_opening(dsn, &gorm.Config{}, WaitOptions{30, 30, 5})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Database connection error: %s", err.Error())
		return
	}

	var help Help
	db.First(&help)
	fmt.Println(help.Description)

	productController := controller.ProductController{}

	r := gin.Default()
	r.GET("/ping", ping)

	r.GET("/v1/get/beans", productController.GetCoffeeBeans)
	r.GET("/v1/get/bean/:id", productController.GetCoffeeBeansById)
	r.POST("/v1/post/bean", productController.PostCoffeeBeans)

	r.GET("/v1/get/cups", productController.GetCups)
	r.GET("/v1/get/cup/:id", productController.GetCupsById)
	r.POST("/v1/post/cup", productController.PostCup)
	r.Run()
}
