package loggermiddleware

import (
	"net/http"

	"github.com/psmey/time_tracker/internal/http/api"
	"github.com/psmey/time_tracker/internal/logger"
)

func New() api.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			logger.LogInfo("Handling Request", "request", request)

			next.ServeHTTP(writer, request)
		})
	}
}
