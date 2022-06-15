package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func (handler handler) DeleteUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("delete_user_handler")

	handler.Logger.Info(Namespace.Concat("DeleteUserHandler"), "")

	comm := communication.New()
	return comm.Response(200, "success", "")
}
