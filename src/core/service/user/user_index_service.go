package user

import (
	"context"

	"github.com/kristiansantos/learning/src/core/domain"
	"github.com/kristiansantos/learning/src/core/repository"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
	"go.mongodb.org/mongo-driver/bson"
)

// var Namespace = namespace.New("service.user.user_index_service")

type IndexUserService struct {
	Context    context.Context
	Repository repository.IUserRepository
	Logger     logger.ILoggerProvider
}

type IIndexUserService interface {
	IndexUser(filter bson.M)
}

func (service *IndexUserService) IndexUser(filter bson.M) (domain.Users, error) {
	return service.Repository.List(filter)
}
