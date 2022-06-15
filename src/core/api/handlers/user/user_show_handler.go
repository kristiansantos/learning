package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/core/api/presenter"
	httphelper "github.com/kristiansantos/learning/src/shared/provider/http_helper"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
	"github.com/kristiansantos/learning/src/shared/tools/parse"
	"go.mongodb.org/mongo-driver/mongo"
)

func (handler handler) ShowUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("show_user_handler")

	handler.Logger.Info(Namespace.Concat("ShowUserHandler"), "")

	var userPresenter presenter.UserShow

	ctx := r.Context()
	comm := communication.New()
	service := handler.Service(ctx).Show
	id := httphelper.GetParam(r, "id")

	switch user, err := service.ShowUser(id); err {
	case nil:
		parse.Unmarshal(user, &userPresenter)
		return comm.Response(200, "success", userPresenter)
	case mongo.ErrNoDocuments:
		service.Logger.Error(Namespace.Concat("Execute", "GetById"), err.Error())
		return comm.ResponseError(404, "not_found", err)
	default:
		service.Logger.Error(Namespace.Concat("Execute", "GetById"), err.Error())
		return comm.ResponseError(400, "error_list", err)
	}
}
