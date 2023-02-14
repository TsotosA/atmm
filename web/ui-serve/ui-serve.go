package ui_serve

import (
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tsotosa/atmm/config"
	"io/fs"
)

//go:embed web/ui-react/build/*
var static embed.FS

func Up() {
	e := echo.New()
	_, _ = fs.Sub(static, "web/ui-react/build")
	e.Static("/", "web/ui-react/build")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Conf.UiPort)))
}
