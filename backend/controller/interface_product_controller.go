package controller

import "github.com/gin-gonic/gin"

type IProductController interface {
	GetProducts(*gin.Context)
	GetProductById(c *gin.Context)
	PostProduct(c *gin.Context)
}
