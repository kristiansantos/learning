package routes

import (
	"github.com/go-chi/chi"
	"github.com/kristiansantos/learning/src/core/api/handlers"
	"github.com/kristiansantos/learning/src/core/api/routes/placement"
	"github.com/kristiansantos/learning/src/shared/middlewares"
)

type router struct {
	Client   *chi.Mux
	Handlers handlers.IHandler
}

func NewRoutes(handlers handlers.IHandler) *router {
	return &router{
		Client:   chi.NewRouter(),
		Handlers: handlers,
	}
}

func (r *router) Setup() {
	// middlewares start setup
	middlewares.LoadOptions(r.Client)
	middlewares.MethodNotAllowed(r.Client)
	middlewares.NotFound(r.Client)
	middlewares.ResourceStatus(r.Client)

	// set routes
	placement.NewRoutes(r.Client, r.Handlers.NewPlacementHandler())
}
