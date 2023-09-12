package core_services

import (
	"log"

	guuid "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/abusquets/ab-kyc/internal/app/errors"
	core_entities "github.com/abusquets/ab-kyc/internal/core/domain/entities"
	core_ports "github.com/abusquets/ab-kyc/internal/core/domain/ports/repositories"
)

// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func GeneratePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

type UserService interface {
	GetByID(uuid string) (*core_entities.User, *errors.AppError)
	CreateUser(inData core_ports.CreateUserInDTO) (user *core_entities.User, error *errors.AppError)
}

type userService struct {
	userRepository core_ports.IUserRepository
}

func NewUserService(userRepository core_ports.IUserRepository) UserService {
	return userService{
		userRepository: userRepository,
	}
}

func (us userService) GetByID(uuid string) (*core_entities.User, *errors.AppError) {
	userUuid, err := guuid.Parse(uuid)
	if err != nil {
		return nil, errors.NewValidationError("Invalid UUID")
	}
	return us.userRepository.GetByID(userUuid.String())
}

func (us userService) CreateUser(inData core_ports.CreateUserInDTO) (user *core_entities.User, error *errors.AppError) {
	if inData.Password != nil {
		fixed := GeneratePassword([]byte(*inData.Password))
		inData.Password = &fixed
	}
	return us.userRepository.Create(inData)
}
