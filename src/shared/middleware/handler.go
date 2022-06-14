package middleware

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/provider/logger"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func Handler(handler func(r *http.Request) communication.Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Namespace.AddComponent("handler.Handler")

		log := logger.Instance
		response := handler(r)
		if response.Error != nil {
			log.Error(Namespace.Concat("handler_error"), response.Error.Error())
		}

		if err := response.JSON(w); err != nil {
			log.Error(Namespace.Concat("response_error"), response.Error.Error())
		}
	}
}
