package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func (handler handler) CreateUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("create_user_handler")

	handler.Logger.Info(Namespace.Concat("CreateUserHandler"), "")

	comm := communication.New()
	return comm.Response(200, "success", "")
}
