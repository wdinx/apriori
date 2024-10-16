package controller

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"apriori-backend/service"
	"apriori-backend/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AprioriControllerImpl struct {
	aprioriService    service.AprioriService
	productRepository repository.ProductRepository
}

func NewAprioriController(aprioriService service.AprioriService, productRepository repository.ProductRepository) AprioriController {
	return &AprioriControllerImpl{aprioriService: aprioriService, productRepository: productRepository}
}

func (controller *AprioriControllerImpl) Apriori(c echo.Context) error {
	var apriori web.CreateAprioriRequest
	c.Bind(&apriori)

	result, err := controller.aprioriService.ProcessApriori(&apriori)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Apriori Success", result))
}

func (controller *AprioriControllerImpl) GetAll(c echo.Context) error {
	var err error
	pageParam := c.QueryParam("page")
	metadata := util.GetMetadata(pageParam)
	metadata.TotalItem, err = controller.productRepository.GetTotalPage(&domain.AprioriResult{}, metadata.Limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	result, err := controller.aprioriService.GetAll(metadata)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessPaginationResponse("Success", *metadata, result))
}

func (controller *AprioriControllerImpl) GetByID(c echo.Context) error {
	strID := c.Param("id")
	result, err := controller.aprioriService.GetByID(strID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", result))
}

func (controller *AprioriControllerImpl) DeleteByID(c echo.Context) error {
	strID := c.Param("id")
	err := controller.aprioriService.DeleteByID(strID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", nil))
}
