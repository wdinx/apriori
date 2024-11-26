package main

import (
	"apriori-backend/config"
	"apriori-backend/route"
	"apriori-backend/util"
	"github.com/go-co-op/gocron/v2"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {

	util.LoadEnv()
	cfg := config.Get()
	e := echo.New()
	validate := validator.New()
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	db := config.InitDB(cfg.Database)
	route.InitRoute(db, e, validate, cfg, scheduler)

	err = e.Start(":3007")
	if err != nil {
		panic(err)
	}
}
