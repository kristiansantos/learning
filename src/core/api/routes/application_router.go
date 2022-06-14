package routes

import (
	"github.com/go-chi/chi"
	"github.com/kristiansantos/learning/src/core/api/handlers"
	"github.com/kristiansantos/learning/src/core/api/routes/user"
	"github.com/kristiansantos/learning/src/shared/middleware"
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
	middleware.LoadOptions(r.Client)
	middleware.MethodNotAllowed(r.Client)
	middleware.NotFound(r.Client)
	middleware.ResourceStatus(r.Client)

	// set routes
	user.NewRoutes(r.Client, r.Handlers.NewUserHandler())
}
