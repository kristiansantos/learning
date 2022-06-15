package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func (handler handler) UpdateUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("update_user_handler")

	handler.Logger.Info(Namespace.Concat("UpdateUserHandler"), "")

	comm := communication.New()
	return comm.Response(200, "success", "")
}
