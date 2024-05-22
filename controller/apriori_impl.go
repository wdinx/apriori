package controller

import (
	"apriori-backend/model/web"
	"apriori-backend/service"
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

	result, err := controller.aprioriService.GetApriori(&apriori)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.NewBaseErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, web.NewBaseSuccessResponse("Apriori Success", result))
}
