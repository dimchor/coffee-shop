package controller

import (
	"coffee_shop_backend/service"
	"coffee_shop_backend/types"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const MAX_AGE = 604800 // one week in seconds
const FRONTEND_IP_PORT = "http://127.0.0.1:4200"
const SESSION_COOKIE = "session_id"

type ProductController struct {
	productService service.IProductService
}

func setupHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", FRONTEND_IP_PORT)
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
	setupHeader(c)
	products, err := p.productService.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var result []types.ProductGetDto
	for _, product := range products {
		result = append(result, *product.ToProductGetDto())
	}
	c.JSON(http.StatusOK, result)
}

func (p *ProductController) GetProductById(c *gin.Context) {
	setupHeader(c)

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
		c.JSON(http.StatusOK, result.ToProductGetDto())
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

func (p *ProductController) PostProduct(c *gin.Context) {
	setupHeader(c)
	var productDto types.ProductPostDto

	if err := c.BindJSON(&productDto); err != nil {
		// status code 400 should be ok
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !p.hasAdminRights(c) {
		c.JSON(http.StatusForbidden, errors.New("nuh uh!"))
		return
	}

	product := productDto.ToProduct()

	if err := p.productService.PostProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (p *ProductController) PostNewUser(c *gin.Context) {
	setupHeader(c)

	var userDto types.UserCreateDto

	if err := c.BindJSON(&userDto); err != nil {
		// status code 400 should be ok
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if userDto.AdminRights && !p.hasAdminRights(c) {
		c.JSON(http.StatusForbidden, errors.New("nuh uh!"))
		return
	}

	if err := p.productService.PostNewUser(&userDto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, userDto)
}

func (p *ProductController) PostLoginUser(c *gin.Context) {
	setupHeader(c)
	var userDto types.UserLoginDto

	if err := c.BindJSON(&userDto); err != nil {
		// status code 400 should be ok
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := p.productService.PostLoginUser(&userDto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.SetCookie(SESSION_COOKIE, token, MAX_AGE, "", "localhost", false, true)
	c.JSON(http.StatusOK, userDto)
}

func (p *ProductController) PostLogoutUser(c *gin.Context) {
	setupHeader(c)
	c.SetCookie(SESSION_COOKIE, "", -1, "", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func (p *ProductController) hasAdminRights(c *gin.Context) bool {
	cookie, err := c.Cookie(SESSION_COOKIE)
	if err != nil {
		return false
	}
	return p.productService.HasAdminRights(cookie)
}

func (p *ProductController) HasAdminRights(c *gin.Context) {
	setupHeader(c)

	if p.hasAdminRights(c) {
		c.JSON(http.StatusForbidden, gin.H{"message": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": true})
}
