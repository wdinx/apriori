package main

import (
	"apriori-backend/config"
	"apriori-backend/route"
	"apriori-backend/util"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {

	util.LoadEnv()
	cfg := config.Get()
	e := echo.New()
	validate := validator.New()

	db := config.InitDB(cfg.Database)
	route.InitRoute(db, e, validate, cfg)

	err := e.Start(":3000")
	if err != nil {
		panic(err)
	}
}
