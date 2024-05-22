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
	transactionRepository := repository.NewTransactionRepository(db)

	userService := service.NewUserService(userRepository, validate)
	imageService := service.NewImageService(imageRepository)
	productService := service.NewProductService(productRepository, imageService, validate)
	transactionService := service.NewTransactionService(transactionRepository, validate)
	aprioriService := service.NewAprioriService(transactionRepository, validate)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService)
	transactionController := controller.NewTransactionController(transactionService)
	aprioriController := controller.NewAprioriController(aprioriService)

	e.POST("/login", userController.Login)
	e.POST("/register", userController.Register)

	router := e.Group("user")
	router.Use(echojwt.JWT([]byte(constant.SECRET_JWT)))

	// Route
	router.GET("/products", productController.GetAll)
	router.GET("/product/:id", productController.GetByID)
	router.POST("/product", productController.Create)
	router.PUT("/product/:id", productController.Update)
	router.DELETE("/product/:id", productController.Delete)

	// Transaction
	router.GET("/transactions", transactionController.GetAll)
	router.GET("/transaction/:id", transactionController.GetById)
	router.POST("/transaction", transactionController.Create)
	router.PUT("/transaction/:id", transactionController.Update)
	router.DELETE("/transaction/:id", transactionController.Delete)

	// Apriori
	router.POST("/apriori", aprioriController.Apriori)
}
