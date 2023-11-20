package main

import (
	"coffee_shop_backend/ping"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello, world")

	r := gin.Default()
	r.GET("/ping", ping.Pong)
	r.Run()
}
