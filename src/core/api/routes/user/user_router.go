package user

import (
	"github.com/go-chi/chi"

	"github.com/kristiansantos/learning/src/core/api/handlers/user"
	middlewares "github.com/kristiansantos/learning/src/shared/middleware"
)

func NewRoutes(router *chi.Mux, userHandler user.IHandler) {
	router.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", middlewares.Handler(userHandler.IndexUserHandler))
	})
}
