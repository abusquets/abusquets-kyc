package core_cli

import (
	core_entities "github.com/abusquets/ab-kyc/internal/core/domain/entities"
)

type CreateUserResponse struct {
	Uuid     string
	Email    string
	IsActive bool
	IsAdmin  bool
}

type UserResponse struct {
	Uuid     string
	Email    string
	IsActive bool
	IsAdmin  bool
}

type UserPresenter interface {
	Present(data *core_entities.User)
	Result() *UserResponse
}

type userPresenter struct {
	UserPresenter
	result *UserResponse
}

func (p *userPresenter) Present(data *core_entities.User) {
	if data == nil {
		return
	}
	p.result = &UserResponse{
		Uuid:  data.Uuid.String(),
		Email: data.Email,
	}
}

func (p *userPresenter) Result() *UserResponse {
	return p.result
}

func NewCreateUserPresenter() UserPresenter {
	return &userPresenter{}
}
