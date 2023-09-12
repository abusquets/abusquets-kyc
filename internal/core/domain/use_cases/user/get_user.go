package core_user_use_cases

import (
	guuid "github.com/google/uuid"

	"github.com/abusquets/ab-kyc/internal/app/errors"
	core_entities "github.com/abusquets/ab-kyc/internal/core/domain/entities"
	core_services "github.com/abusquets/ab-kyc/internal/core/domain/services"
	ipresenter "github.com/abusquets/ab-kyc/internal/shared/presenter"
)

type GetUserUseCase interface {
	Execute(uuid string) *errors.AppError
}

type getUserUseCase struct {
	userService core_services.UserService
	presenter   ipresenter.IPresenterIn[core_entities.User]
}

func NewGetUserUseCase(userService core_services.UserService, p ipresenter.IPresenterIn[core_entities.User]) GetUserUseCase {
	return &getUserUseCase{
		userService: userService,
		presenter:   p,
	}
}

func (uc *getUserUseCase) Execute(uuid string) *errors.AppError {
	userUuid, err := guuid.Parse(uuid)
	if err != nil {
		return errors.NewValidationError("Invalid UUID")
	}
	user, error := uc.userService.GetByID(userUuid.String())

	if error != nil {
		return error
	}

	uc.presenter.Present(user)

	return nil
}
