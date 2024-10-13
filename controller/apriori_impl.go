package controller

import (
	"apriori-backend/model/web"
	"apriori-backend/service"
	"apriori-backend/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AprioriControllerImpl struct {
	aprioriService service.AprioriService
}

func NewAprioriController(aprioriService service.AprioriService) AprioriController {
	return &AprioriControllerImpl{aprioriService: aprioriService}
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
	pageParam := c.QueryParam("page")
	metadata := util.GetMetadata(pageParam)
	result, err := controller.aprioriService.GetAll(metadata)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Success", result))
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
