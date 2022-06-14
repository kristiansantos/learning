package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func (handler handler) IndexUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("index_user_handler")

	handler.Logger.Info(Namespace.Concat("IndexUserHandler"), "")

	comm := communication.New()
	return comm.Response(200, "success", "")
}
