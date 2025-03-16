package controller

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"apriori-backend/service"
	"apriori-backend/util"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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

	dateStart, err := time.Parse("2006-01-02", apriori.DateStart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse("Invalid Date Start Format"))
	}
	dateEnd, err := time.Parse("2006-01-02", apriori.DateEnd)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse("Invalid Date End Format"))
	}
	if dateStart.After(dateEnd) {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse("Tanggal mulai tidak boleh lebih besar dari tanggal selesai"))
	}

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
	metadata.TotalItem, err = controller.productRepository.GetTotalPage(&domain.AprioriData{}, metadata.Limit)
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
	data := result.GetRecommendation(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", data))
}

func (controller *AprioriControllerImpl) DeleteByID(c echo.Context) error {
	strID := c.Param("id")
	err := controller.aprioriService.DeleteByID(strID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", nil))
}

func (controller *AprioriControllerImpl) GetRecommendationItem(c echo.Context) error {
	result, err := controller.aprioriService.GetRecommendationItem()
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", result))
}

func (controller *AprioriControllerImpl) CreateRecommendationItem(c echo.Context) error {
	err := controller.aprioriService.CreateRecommendationItem()
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", nil))
}

func (controller *AprioriControllerImpl) DeleteAll(c echo.Context) error {
	err := controller.aprioriService.DeleteAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", nil))
}
