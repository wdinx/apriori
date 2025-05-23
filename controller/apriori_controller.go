package controller

import "github.com/labstack/echo/v4"

type AprioriController interface {
	Apriori(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	DeleteByID(c echo.Context) error
	GetRecommendationItem(c echo.Context) error
	CreateRecommendationItem(c echo.Context) error
	DeleteAll(c echo.Context) error
}
