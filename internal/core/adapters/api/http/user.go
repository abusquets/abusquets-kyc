package core_http

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	core_repositories "github.com/abusquets/ab-kyc/internal/core/adapters/spi/repositories"
	core_services "github.com/abusquets/ab-kyc/internal/core/domain/services"
	core_user_use_cases "github.com/abusquets/ab-kyc/internal/core/domain/use_cases/user"
)

type UserHandlers struct {
	db *sqlx.DB
}

func newUserHandlers(e *echo.Echo, db *sqlx.DB) {
	handler := &UserHandlers{
		db: db,
	}
	e.GET("/core/users/:uuid", handler.getByID)
}

func (uh UserHandlers) getByID(c echo.Context) error {
	userRepo := core_repositories.NewUserRepository(uh.db)
	userService := core_services.NewUserService(userRepo)
	presenter := NewUserPresenter()
	getUserUseCase := core_user_use_cases.NewGetUserUseCase(userService, presenter)

	uuid := c.Param("uuid")

	apperror := getUserUseCase.Execute(uuid)

	if apperror != nil {
		return c.JSON(apperror.Code, apperror.AsMessage())
	}

	return c.JSON(http.StatusOK, presenter.Result())
}
