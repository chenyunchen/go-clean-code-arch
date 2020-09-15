package http

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/config"
)

type Application struct {
	config config.Config
	router *echo.Echo
}

func newApplication(
	config config.Config,
	router *echo.Echo,
) *Application {
	return &Application{
		config: config,
		router: router,
	}
}

func (app *Application) Launch() {
	log.Fatal(app.router.Start(fmt.Sprintf(":%d", app.config.Server.Port)))
}
