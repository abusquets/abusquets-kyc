package core_user_use_cases

import (
	"github.com/abusquets/ab-kyc/internal/app/errors"
	core_entities "github.com/abusquets/ab-kyc/internal/core/domain/entities"
	core_ports "github.com/abusquets/ab-kyc/internal/core/domain/ports/repositories"
	core_services "github.com/abusquets/ab-kyc/internal/core/domain/services"
	ipresenter "github.com/abusquets/ab-kyc/internal/shared/presenter"
)

type CreateUserUseCase interface {
	Execute(inData core_ports.CreateUserInDTO) *errors.AppError
}

type createUserUseCase struct {
	userService core_services.UserService
	presenter   ipresenter.IPresenterIn[core_entities.User]
}

func NewCreateUserUseCase(userService core_services.UserService, p ipresenter.IPresenterIn[core_entities.User]) CreateUserUseCase {
	return &createUserUseCase{
		userService: userService,
		presenter:   p,
	}
}

func (uc *createUserUseCase) Execute(inData core_ports.CreateUserInDTO) *errors.AppError {
	user, error := uc.userService.CreateUser(inData)
	if error != nil {
		return error
	}
	uc.presenter.Present(user)
	return nil
}
