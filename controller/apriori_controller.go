package controller

import "github.com/labstack/echo/v4"

type AprioriController interface {
	Apriori(c echo.Context) error
}
