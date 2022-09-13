package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
	// Put the service
}

// NewAuthController used to create new instance of AuthController
func NewAuthController() AuthController {
	return &authController{}
}

func (c *authController) Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Login OK",
	})
}

func (c *authController) Register(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Register OK",
	})
}
