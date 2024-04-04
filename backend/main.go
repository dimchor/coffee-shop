package main

import (
	"coffee_shop_backend/controller"
	"coffee_shop_backend/service"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ping(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:4200")

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

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
	password := os.Getenv("DB_ROOT_PASSWORD")
	if password == "" {
		fmt.Printf("DB_ROOT_PASSWORD variable not set")
		return
	}

	const database_name = "coffeeshop"
	dsn := "root:" + password + "@tcp(db:3306)/" +
		database_name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := try_opening(dsn, &gorm.Config{}, WaitOptions{30, 30, 10})
	if err != nil {
		fmt.Printf("database connection error: %s\n", err.Error())
		return
	}

	productService, err := service.NewProductService(db)
	if err != nil {
		fmt.Printf("product service error: %s\n", err.Error())
	}

	var productController controller.IProductController = controller.
		NewProductController(productService)

	r := gin.Default()
	r.GET("/ping", ping)

	r.GET("/v1/get/products", productController.GetProducts)
	r.GET("/v1/get/product/:id", productController.GetProductById)
	r.GET("/v1/get/has_admin_rights/:token", productController.HasAdminRights)
	r.GET("/v1/get/user_details/:token", productController.GetUserDetails)
	r.POST("/v1/post/product", productController.PostProduct)
	r.POST("/v1/post/new_user", productController.PostNewUser)
	r.POST("/v1/post/login_user", productController.PostLoginUser)
	r.POST("/v1/post/logout_user", productController.PostLogoutUser)

	r.Run()
}
