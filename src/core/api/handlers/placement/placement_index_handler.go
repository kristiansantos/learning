package placement

import (
	"net/http"

	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func (handler handler) IndexPlacementHandler(r *http.Request) communication.Response {
	Namespace.AddComponent("index_placement_handler")

	handler.Logger.Info(Namespace.Concat("IndexPlacementHandler"), "")

	comm := communication.New()
	return comm.Response(200, "success", "")
}
