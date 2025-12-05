package authmiddleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/psmey/time_tracker/internal/http/api"
	"github.com/psmey/time_tracker/internal/logger"
)

type AuthProvider interface {
	ValidateToken(context context.Context, token string) (*jwt.MapClaims, error)
}

func New(authProvider AuthProvider) api.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			logger.LogDebug("Handling authentication for request", "request", request)
			authHeader := request.Header.Get("Authorization")

			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				logger.LogWarn("Invalid authorization header", "request", request)
				http.Error(writer, "Unauthorized", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := authProvider.ValidateToken(request.Context(), token)
			if err != nil {
				logger.LogWarn("Invalid token used", "request", request)
				http.Error(writer, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claimsContext := context.WithValue(request.Context(), "claims", claims)
			next.ServeHTTP(writer, request.WithContext(claimsContext))
		})
	}
}
