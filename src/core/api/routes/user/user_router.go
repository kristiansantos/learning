package user

import (
	"github.com/go-chi/chi"

	"github.com/kristiansantos/learning/src/core/api/handlers/user"
	"github.com/kristiansantos/learning/src/shared/middlewares"
)

func NewRoutes(router *chi.Mux, userHandler user.IHandler) {
	router.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", middlewares.Handler(userHandler.IndexUserHandler))
	})
}
