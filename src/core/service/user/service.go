package user

import (
	"context"

	"github.com/kristiansantos/learning/src/core/repository"
	"github.com/kristiansantos/learning/src/shared/provider/hash"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
)

type Dependencies struct {
	Context    context.Context
	Repository repository.IUserRepository
	Logger     logger.ILoggerProvider
	Hash       hash.IHashProvider
}

type Services struct {
	Index  IndexUserService
	Show   ShowUserService
	Create CreateUserService
	Update UpdateUserService
	Delete DeleteUserService
}

func NewUser(dep Dependencies) *Services {
	return &Services{
		Index: IndexUserService{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Show: ShowUserService{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Create: CreateUserService{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Update: UpdateUserService{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
		Delete: DeleteUserService{
			Repository: dep.Repository,
			Logger:     dep.Logger,
		},
	}
}
