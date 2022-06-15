package handlers

import (
	"github.com/kristiansantos/learning/src/core/api/handlers/user"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
)

type Dependencies struct {
	Logger logger.ILoggerProvider
}

type IHandler interface {
	NewUserHandler() user.IHandler
}
type handler struct {
	userHandler user.IHandler
}

func NewHandler(dep Dependencies) handler {
	return handler{
		userHandler: user.NewHandler(dep.Logger),
	}
}

func (h handler) NewUserHandler() user.IHandler {
	return h.userHandler
}
