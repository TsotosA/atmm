package ui_serve

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tsotosa/atmm/config"
	"io/fs"
)

//go:embed build/*
var static embed.FS

func Up() {
	e := echo.New()
	_, _ = fs.Sub(static, "build")
	e.Static("/", "build")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Conf.UiPort)))
}
