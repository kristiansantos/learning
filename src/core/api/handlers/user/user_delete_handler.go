package user

import (
	"net/http"

	httphelper "github.com/kristiansantos/learning/src/shared/provider/http_helper"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func (handler handler) DeleteUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("delete_user_handler")

	handler.Logger.Info(Namespace.Concat("DeleteUserHandler"), "")

	ctx := r.Context()
	comm := communication.New()
	service := handler.Service(ctx).Delete
	id := httphelper.GetParam(r, "id")

	service.DeleteUser(id)

	return comm.Response(200, "success", "")
}
