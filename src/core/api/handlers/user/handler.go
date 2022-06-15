package user

import (
	"context"
	"net/http"

	userService "github.com/kristiansantos/learning/src/core/service/user"
	userRepo "github.com/kristiansantos/learning/src/infrastructure/mongodb/user"
	"github.com/kristiansantos/learning/src/shared/provider/hash"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
	"github.com/kristiansantos/learning/src/shared/tools/namespace"
)

var Namespace = namespace.New("core.api.handlers.user")

type IHandler interface {
	CreateUserHandler(r *http.Request) communication.Response
	DeleteUserHandler(r *http.Request) communication.Response
	IndexUserHandler(r *http.Request) communication.Response
	ShowUserHandler(r *http.Request) communication.Response
	UpdateUserHandler(r *http.Request) communication.Response
}

type handler struct {
	Logger logger.ILoggerProvider
}

func NewHandler(logger logger.ILoggerProvider) handler {
	return handler{
		Logger: logger,
	}
}

func (a handler) Service(ctx context.Context) *userService.Services {
	userRepository := userRepo.Setup(ctx)

	dependencies := userService.Dependencies{
		Context:    ctx,
		Repository: userRepository,
		Logger:     logger.New(),
		Hash:       hash.New(),
	}
	return userService.NewUser(dependencies)
}
