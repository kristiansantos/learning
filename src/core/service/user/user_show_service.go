package user

import (
	"context"

	"github.com/kristiansantos/learning/src/core/domain"
	"github.com/kristiansantos/learning/src/core/repository"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
)

// var Namespace = namespace.New("service.user.create_service")

type ShowUserService struct {
	Context    context.Context
	Repository repository.IUserRepository
	Logger     logger.ILoggerProvider
}

type IShowUserService interface {
	ShowUser(id string)
}

func (service *ShowUserService) ShowUser(id string) (domain.User, error) {
	return service.Repository.GetBy(id)
}
