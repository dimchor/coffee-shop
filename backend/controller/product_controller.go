package controller

import (
	"coffee_shop_backend/service"
	"coffee_shop_backend/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	// temporary, GORM will be used in the (near) future
	/*
		coffeebeans []product.CoffeeBean
		cups        []product.Cup
	*/
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	/*
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
	*/
	return &ProductController{productService}
}

func (p *ProductController) GetProducts(c *gin.Context) {
	result, err := p.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (p *ProductController) GetProductById(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := p.productService.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if result != nil {
		c.JSON(http.StatusOK, result)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "coffee bean not found"})
}

func (p *ProductController) PostProduct(c *gin.Context) {
	var product types.Product

	if err := c.BindJSON(&product); err != nil {
		// status code 400 should be ok
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := p.productService.PostProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}
