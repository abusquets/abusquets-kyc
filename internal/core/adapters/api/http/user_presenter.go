package core_http

import (
	core_entities "github.com/abusquets/ab-kyc/internal/core/domain/entities"
	ipresenter "github.com/abusquets/ab-kyc/internal/shared/presenter"
)

type UserResponse struct {
	Uuid  string `json:"uuid"`
	Email string `json:"email"`
}

type UserPresenter interface {
	ipresenter.IPresenterIn[core_entities.User]
	ipresenter.IPresenterOut[UserResponse]
}

type userPresenter struct {
	result *UserResponse
}

func (p *userPresenter) Present(data *core_entities.User) {
	if data != nil {
		p.result = &UserResponse{
			Uuid:  data.Uuid.String(),
			Email: data.Email,
		}
	}
}

func (p *userPresenter) Result() *UserResponse {
	return p.result
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}
