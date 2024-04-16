package controller

import "github.com/gin-gonic/gin"

type IProductController interface {
	GetProducts(*gin.Context)
	GetProductById(*gin.Context)
	GetUserDetails(*gin.Context)
	PostProduct(*gin.Context)
	PostNewUser(*gin.Context)
	PostLoginUser(*gin.Context)
	PostLogoutUser(*gin.Context)
	HasAdminRights(*gin.Context)
	PostNewPassword(c *gin.Context)
}
