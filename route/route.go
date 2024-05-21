package route

import (
	"apriori-backend/config"
	"apriori-backend/constant"
	"apriori-backend/controller"
	"apriori-backend/repository"
	"apriori-backend/service"
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB, e *echo.Echo, validate *validator.Validate, config *config.Config) {
	// Your code here
	userRepository := repository.NewUserRepository(db)
	imageRepository := repository.NewImageRepository(config.DigitalOceanSpaces)
	productRepository := repository.NewProductRepository(db)

	userService := service.NewUserService(userRepository, validate)
	imageService := service.NewImageService(imageRepository)
	productService := service.NewProductService(productRepository, imageService, validate)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)

	e.POST("/login", userController.Login)
	e.POST("/register", userController.Register)

	product := e.Group("user")
	product.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))
	product.GET("/products", productController.GetAll)
	product.GET("/product/:id", productController.GetByID)
	product.POST("/product", productController.Create)
	product.PUT("/product/:id", productController.Update)
	product.DELETE("/product/:id", productController.Delete)
}
