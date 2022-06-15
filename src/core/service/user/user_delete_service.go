package user

import (
	"context"

	"github.com/kristiansantos/learning/src/core/repository"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
)

// var Namespace = namespace.New("service.user.create_service")

type DeleteUserService struct {
	Context    context.Context
	Repository repository.IUserRepository
	Logger     logger.ILoggerProvider
}

type IDeleteUserService interface {
	DeleteUser(id string)
}

func (service *DeleteUserService) DeleteUser(id string) error {
	return service.Repository.Delete(id)
}
