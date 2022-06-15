package user

import (
	"context"

	"github.com/kristiansantos/learning/src/core/domain"
	"github.com/kristiansantos/learning/src/core/repository"
	"github.com/kristiansantos/learning/src/core/service/dto"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
	"github.com/kristiansantos/learning/src/shared/tools/parse"
)

// var Namespace = namespace.New("service.user.create_service")

type CreateUserService struct {
	Context    context.Context
	Repository repository.IUserRepository
	Logger     logger.ILoggerProvider
}

type ICreateUserService interface {
	CreateUser(dto dto.UserDTO) (interface{}, error)
}

func (service *CreateUserService) CreateUser(dto dto.UserDTO) (interface{}, error) {
	var user domain.User

	parse.Unmarshal(dto, &user)
	user.Populate()
	return service.Repository.Create(user)
}
