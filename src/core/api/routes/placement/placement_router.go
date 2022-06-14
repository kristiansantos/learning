package placement

import (
	"github.com/go-chi/chi"

	"github.com/kristiansantos/learning/src/core/api/handlers/placement"
	"github.com/kristiansantos/learning/src/shared/middlewares"
)

func NewRoutes(router *chi.Mux, placementHandler placement.IHandler) {
	router.Route("/api/v1/placements", func(r chi.Router) {
		r.Get("/", middlewares.Handler(placementHandler.IndexPlacementHandler))
	})
}
