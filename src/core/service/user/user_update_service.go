package user

import (
	"context"
	"time"

	"github.com/kristiansantos/learning/src/core/domain"
	"github.com/kristiansantos/learning/src/core/repository"
	"github.com/kristiansantos/learning/src/core/service/dto"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
	"github.com/kristiansantos/learning/src/shared/tools/parse"
)

// var Namespace = namespace.New("service.user.create_service")

type UpdateUserService struct {
	Context    context.Context
	Repository repository.IUserRepository
	Logger     logger.ILoggerProvider
}

type IUpdateUserService interface {
	UpdateUser(id string, dto dto.UserDTO) (interface{}, error)
}

func (service *UpdateUserService) UpdateUser(id string, dto dto.UserDTO) (interface{}, error) {
	var user domain.User

	parse.Unmarshal(dto, &user)

	user.ID = id
	user.UpdatedAt = time.Now()

	return service.Repository.Update(id, user)
}
