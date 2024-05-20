package controller

import (
	"apriori-backend/model/web"
	"apriori-backend/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{userService: userService}
}

func (controller *UserControllerImpl) Login(c echo.Context) error {
	var user web.LoginRequest
	c.Bind(&user)
	response, err := controller.userService.Login(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Login success", response))
}

func (controller *UserControllerImpl) Register(c echo.Context) error {
	var user web.RegisterRequest
	c.Bind(&user)
	err := controller.userService.Register(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Register success", nil))
}
