package route

import (
	"apriori-backend/config"
	"apriori-backend/constant"
	"apriori-backend/controller"
	"apriori-backend/middleware"
	"apriori-backend/repository"
	"apriori-backend/service"
	"github.com/go-co-op/gocron/v2"
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(db *gorm.DB, e *echo.Echo, validate *validator.Validate, config *config.Config, scheduler gocron.Scheduler) {
	// Your code here
	userRepository := repository.NewUserRepository(db)
	imageRepository := repository.NewImageRepository(config.Cloudinary)
	productRepository := repository.NewProductRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)
	aprioriRepository := repository.NewAprioriRepository(db)
	recommendationRepository := repository.NewRecommendationRepository(db)

	userService := service.NewUserService(userRepository, validate)
	imageService := service.NewImageService(imageRepository)
	productService := service.NewProductService(productRepository, imageService, validate)
	transactionService := service.NewTransactionService(transactionRepository, validate)
	aprioriService := service.NewAprioriService(transactionRepository, aprioriRepository, recommendationRepository, validate)
	gocronService := service.NewCronJobServiceImpl(aprioriService, scheduler)

	userController := controller.NewUserController(userService)
	productController := controller.NewProductController(productService, productRepository)
	transactionController := controller.NewTransactionController(transactionService, productRepository)
	aprioriController := controller.NewAprioriController(aprioriService, productRepository)

	gocronService.InitCronJob()

	e.Use(middleware.CORSMiddleware())

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
	router.POST("/transactions/excel", transactionController.InsertByExcel)
	router.DELETE("/transactions/delete_all", transactionController.DeleteAll)

	// Apriori
	e.POST("/apriori", aprioriController.Apriori)
	e.GET("/apriori", aprioriController.GetAll)
	e.GET("/apriori/:id", aprioriController.GetByID)
	e.DELETE("/apriori/:id", aprioriController.DeleteByID)

	// Recommendation
	e.GET("/recommendations", aprioriController.GetRecommendationItem)
	e.POST("/recommendations", aprioriController.CreateRecommendationItem)
}
