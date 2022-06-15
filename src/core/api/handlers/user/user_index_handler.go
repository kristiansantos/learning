package user

import (
	"net/http"

	"github.com/kristiansantos/learning/src/core/api/presenter"
	"github.com/kristiansantos/learning/src/shared/tools/communication"
	"github.com/kristiansantos/learning/src/shared/tools/parse"
	"go.mongodb.org/mongo-driver/bson"
)

func (handler handler) IndexUserHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("index_user_handler")

	handler.Logger.Info(Namespace.Concat("IndexUserHandler"), "")

	var userSlice presenter.UsersIndex

	ctx := r.Context()
	comm := communication.New()
	service := handler.Service(ctx).Index
	filter := bson.M{}

	users, _ := service.IndexUser(filter)

	for _, user := range users {
		var userPresenter presenter.UserIndex
		parse.Unmarshal(user, &userPresenter)
		userSlice = append(userSlice, &userPresenter)
	}

	return comm.Response(200, "success", userSlice)
}
