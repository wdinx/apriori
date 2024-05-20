package route

import (
	"apriori-backend/config"
	"apriori-backend/controller"
	"apriori-backend/repository"
	"apriori-backend/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB, e *echo.Echo, validate *validator.Validate, config *config.Config) {
	// Your code here
	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository, validate)

	userController := controller.NewUserController(userService)

	e.POST("/login", userController.Login)
	e.POST("/register", userController.Register)
}
