package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kristiansantos/learning/src/shared/tools/communication"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				comm := communication.New()

				fmt.Println(err)

				response := comm.ResponseError(500, "error", errors.New("internal server error"))
				response.JSON(w)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
