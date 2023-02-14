package ui_serve

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tsotosa/atmm/config"
)

func Up() {
	e := echo.New()
	//_, _ = fs.Sub(static, "/web/ui-react/build")
	e.Static("/", "../ui-react/build")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Conf.UiPort)))
}
