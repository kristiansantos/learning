package handlers

import (
	"github.com/kristiansantos/learning/src/core/api/handlers/placement"
	"github.com/kristiansantos/learning/src/shared/providers/logger"
)

// Implementa struct de dependecias
type Dependencies struct {
	Logger logger.ILoggerProvider
}

// Implementa interface de hendler
type IHandler interface {
	NewPlacementHandler() placement.IHandler
}

// Implementa struct de handler
type handler struct {
	PlacementHandler placement.IHandler
}

// Função que retorna uma struct de handler com as dependencias
func NewHandler(dep Dependencies) handler {
	return handler{
		PlacementHandler: placement.NewHandler(dep.Logger),
	}
}

// Implementa metodo declarado na struct
func (h handler) NewPlacementHandler() placement.IHandler {
	return h.PlacementHandler
}
