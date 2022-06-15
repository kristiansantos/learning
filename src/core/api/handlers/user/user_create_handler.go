package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/core/api/presenter"
	"github.com/kristiansantos/learning/src/core/service/dto"
	httphelper "github.com/kristiansantos/learning/src/shared/provider/http_helper"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
	"github.com/kristiansantos/learning/src/shared/tools/parse"
)

func (handler handler) CreateUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("create_user_handler")

	handler.Logger.Info(Namespace.Concat("CreateUserHandler"), "")

	var dto dto.UserDTO
	var userPresenter presenter.UserCreate

	ctx := r.Context()
	comm := communication.New()
	service := handler.Service(ctx).Create
	err := httphelper.GetBody(r.Body, &dto)

	if err != nil {
		handler.Logger.Error(Namespace.Concat("CreatePlanetHandler", "GetBody"), err.Error())

		return comm.ResponseError(400, "error_create", err)
	}

	user, err := service.CreateUser(dto)

	if err != nil {
		handler.Logger.Error(Namespace.Concat("CreatePlanetHandler", "CreateUser"), err.Error())

		return comm.ResponseError(400, "error_create", err)
	}

	parse.Unmarshal(user, &userPresenter)
	return comm.Response(201, "success_create", userPresenter)
}
