package user

import (
	"github.com/go-chi/chi"

	"github.com/kristiansantos/learning/src/core/api/handlers/user"
	middlewares "github.com/kristiansantos/learning/src/shared/middleware"
)

func NewRoutes(router *chi.Mux, userHandler user.IHandler) {
	router.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", middlewares.Handler(userHandler.IndexUserHandler))
		r.Get("/{id}", middlewares.Handler(userHandler.ShowUserHandler))
		r.Delete("/{id}", middlewares.Handler(userHandler.DeleteUserHandler))
		r.Post("/", middlewares.Handler(userHandler.CreateUserHandler))
		r.Put("/{id}", middlewares.Handler(userHandler.UpdateUserHandler))
	})
}
