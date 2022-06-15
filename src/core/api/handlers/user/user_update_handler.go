package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/core/api/presenter"
	"github.com/kristiansantos/learning/src/core/service/dto"
	httphelper "github.com/kristiansantos/learning/src/shared/provider/http_helper"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
	"github.com/kristiansantos/learning/src/shared/tools/parse"
)

func (handler handler) UpdateUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("update_user_handler")

	handler.Logger.Info(Namespace.Concat("UpdateUserHandler"), "")

	var dto dto.UserDTO
	var userPresenter presenter.UserUpdate

	ctx := r.Context()
	comm := communication.New()
	service := handler.Service(ctx).Update
	id := httphelper.GetParam(r, "id")
	err := httphelper.GetBody(r.Body, &dto)

	if err != nil {
		handler.Logger.Error(Namespace.Concat("UpdatePlanetHandler", "GetBody"), err.Error())

		return comm.ResponseError(400, "error_update", err)
	}

	user, err := service.UpdateUser(id, dto)

	if err != nil {
		handler.Logger.Error(Namespace.Concat("UpdatePlanetHandler", "UpdateUser"), err.Error())

		return comm.ResponseError(400, "error_update", err)
	}

	parse.Unmarshal(user, &userPresenter)
	return comm.Response(200, "success", userPresenter)
}
