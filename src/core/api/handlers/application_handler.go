package handlers

import (
	"github.com/kristiansantos/learning/src/core/api/handlers/user"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
)

// Implementa struct de dependecias
type Dependencies struct {
	Logger logger.ILoggerProvider
}

// Implementa interface de hendler
type IHandler interface {
	NewUserHandler() user.IHandler
}

// Implementa struct de handler
type handler struct {
	userHandler user.IHandler
}

// Função que retorna uma struct de handler com as dependencias
func NewHandler(dep Dependencies) handler {
	return handler{
		userHandler: user.NewHandler(dep.Logger),
	}
}

// Implementa metodo declarado na struct
func (h handler) NewUserHandler() user.IHandler {
	return h.userHandler
}
