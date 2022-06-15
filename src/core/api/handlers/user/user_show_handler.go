package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func (handler handler) ShowUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("show_user_handler")

	handler.Logger.Info(Namespace.Concat("ShowUserHandler"), "")

	comm := communication.New()
	return comm.Response(200, "success", "")
}
