package controllers

import (
	"dans-multi-pro/params"
	"dans-multi-pro/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		userService: *service,
	}
}

func (u *UserController) UserRegister(c *gin.Context) {

	var req params.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.CreateUser(req)

	c.JSON(result.Status, result.Payload)
}

func (u *UserController) Login(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.Login(req)

	c.JSON(result.Status, result.Payload)
}
