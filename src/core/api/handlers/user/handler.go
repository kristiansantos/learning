package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/provider/logger"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
	"github.com/kristiansantos/learning/src/shared/tools/namespace"
)

var Namespace = namespace.New("core.api.handlers.user")

type IHandler interface {
	IndexUserHandler(r *http.Request) communication.Response
}

type handler struct {
	Logger logger.ILoggerProvider
}

func NewHandler(logger logger.ILoggerProvider) handler {
	return handler{
		Logger: logger,
	}
}
