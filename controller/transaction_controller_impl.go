package controller

import (
	"apriori-backend/model/web"
	"apriori-backend/service"
	"github.com/labstack/echo/v4"
	"strconv"
)

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionControllerImpl {
	return TransactionControllerImpl{transactionService: transactionService}
}

func (controller *TransactionControllerImpl) Create(c echo.Context) error {
	var createTransactionRequest web.CreateTransactionRequest
	if err := c.Bind(&createTransactionRequest); err != nil {
		return c.JSON(400, web.NewBaseErrorResponse("Bad Request"))
	}
	err := controller.transactionService.Create(&createTransactionRequest)
	if err != nil {
		return c.JSON(500, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(201, web.NewBaseSuccessResponse("Transaction created successfully", nil))
}

func (controller *TransactionControllerImpl) Update(c echo.Context) error {
	var updateTransactionRequest web.UpdateTransactionRequest
	var err error
	if err = c.Bind(&updateTransactionRequest); err != nil {
		return c.JSON(400, web.NewBaseErrorResponse("Bad Request"))
	}
	strID := c.Param("id")
	updateTransactionRequest.ID, err = strconv.Atoi(strID)
	err = controller.transactionService.Update(&updateTransactionRequest)
	if err != nil {
		return c.JSON(500, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(200, web.NewBaseSuccessResponse("Transaction updated successfully", nil))
}

func (controller *TransactionControllerImpl) Delete(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.JSON(400, web.NewBaseErrorResponse("Bad Request"))
	}
	err = controller.transactionService.Delete(id)
	if err != nil {
		return c.JSON(500, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(200, web.NewBaseSuccessResponse("Transaction deleted successfully", nil))
}

func (controller *TransactionControllerImpl) GetById(c echo.Context) error {
	strID := c.Param("id")
	id, err := strconv.Atoi(strID)
	if err != nil {
		return c.JSON(400, web.NewBaseErrorResponse("Bad Request"))
	}
	transaction, err := controller.transactionService.FindById(id)
	if err != nil {
		return c.JSON(500, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(200, web.NewBaseSuccessResponse("Transaction retrieved successfully", transaction))
}

func (controller *TransactionControllerImpl) GetAll(c echo.Context) error {
	transactions, err := controller.transactionService.FindAll()
	if err != nil {
		return c.JSON(500, web.NewBaseErrorResponse(err.Error()))
	}
	return c.JSON(200, web.NewBaseSuccessResponse("Get all transactions success", transactions))
}
