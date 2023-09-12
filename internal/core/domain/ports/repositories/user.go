package core_ports

import (
	"github.com/abusquets/ab-kyc/internal/app/errors"
	core_entities "github.com/abusquets/ab-kyc/internal/core/domain/entities"
)

type CreateUserInDTO struct {
	Email    string
	Password *string
	IsActive bool
	IsAdmin  bool
}

type IUserRepository interface {
	GetByID(uuid string) (*core_entities.User, *errors.AppError)
	Create(userData CreateUserInDTO) (user *core_entities.User, error *errors.AppError)
}
