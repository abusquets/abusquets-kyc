package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	core_http "github.com/abusquets/ab-kyc/internal/core/adapters/api/http"
	"github.com/abusquets/ab-kyc/internal/db"
)

func Start(config *Config, dbManager db.DBManager) {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	core_http.NewCoreRouterHandler(e, dbManager.Database())

	// Start server
	address := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	e.Logger.Fatal(e.Start(address))

}
