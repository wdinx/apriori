package controller

import (
	"apriori-backend/exception"
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"apriori-backend/service"
	"apriori-backend/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	productService    service.ProductService
	productRepository repository.ProductRepository
}

func NewProductController(productService service.ProductService, productRepository repository.ProductRepository) ProductController {
	return &ProductControllerImpl{
		productService:    productService,
		productRepository: productRepository,
	}
}

func (controller *ProductControllerImpl) Create(c echo.Context) error {
	var request web.ProductCreateRequest
	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	request.Image, err = c.FormFile("image")
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	err = controller.productService.Create(&request)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("product created successfully", nil))
}

func (controller *ProductControllerImpl) Update(c echo.Context) error {
	var request web.ProductUpdateRequest

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	request.Image, err = c.FormFile("image")
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}

	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	request.ID = id

	err = controller.productService.Update(&request)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("product updated successfully", nil))
}

func (controller *ProductControllerImpl) Delete(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	err = controller.productService.Delete(id)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("product deleted successfully", nil))
}

func (controller *ProductControllerImpl) GetByID(c echo.Context) error {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	product, err := controller.productService.GetByID(id)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("product retrieved successfully", product))
}

func (controller *ProductControllerImpl) GetAll(c echo.Context) error {
	var err error
	pageParam := c.QueryParam("page")
	metadata := util.GetMetadata(pageParam)
	metadata.TotalData, err = controller.productRepository.GetTotalPage(&domain.Product{}, metadata.Limit)
	products, err := controller.productService.GetAll(metadata)
	if err != nil {
		return c.JSON(exception.ErrorHandler(err), web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessPaginationResponse("products retrieved successfully", *metadata, products))
}
