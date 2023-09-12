package core_http

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type CoreRouterHandler struct {
}

func NewCoreRouterHandler(e *echo.Echo, db *sqlx.DB) {
	newUserHandlers(e, db)
}
