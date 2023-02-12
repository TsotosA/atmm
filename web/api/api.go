package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/helper"
	"github.com/tsotosa/atmm/model"
	"net/http"
	"strconv"
)

func Up() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	api := e.Group("api")
	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api.GET("/log", func(c echo.Context) error {
		nItems := c.QueryParam("nItems")
		grepFor := c.QueryParam("grepFor")
		nLines, err := strconv.Atoi(nItems)
		if err != nil {
		}
		y := ""
		if nLines > 0 {
			y = helper.GetLastNLinesWithSeek(config.Conf.LogOutputPath, nLines)
		}
		if nLines <= 0 {
			y = helper.GetAllFileAsString(config.Conf.LogOutputPath)
		}
		if grepFor != "" {
			grepped := helper.GrepInString(y, grepFor)
			y = grepped
		}
		return c.String(http.StatusOK, y)
	})
	api.GET("/config", func(c echo.Context) error {
		res := config.Conf.Mask()
		return c.JSON(http.StatusOK, res)
	})
	api.POST("/config", func(c echo.Context) error {
		//todo: WIP config edit
		fmt.Println("WIP post")
		conf := new(model.AppConfUpdate)
		if err := c.Bind(conf); err != nil {
			return err
		}
		config.Conf.UpdateFields(conf)
		return c.JSON(http.StatusCreated, conf)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Conf.ApiPort)))
}
