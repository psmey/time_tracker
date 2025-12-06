package logger

import "net/http"

func LogHttpDebug(message string, request *http.Request, args ...any) {
	allArgs := buildWithHttpArgs(request, args)
	LogDebug(message, allArgs)
}

func LogHttpInfo(message string, request *http.Request, args ...any) {
	allArgs := buildWithHttpArgs(request, args)
	LogInfo(message, allArgs)
}

func LogHttpWarn(message string, request *http.Request, args ...any) {
	allArgs := buildWithHttpArgs(request, args)
	LogWarn(message, allArgs)
}

func LogHttpError(message string, err error, request *http.Request, args ...any) {
	allArgs := buildWithHttpArgs(request, args)
	LogError(message, err, allArgs...)
}

func buildWithHttpArgs(request *http.Request, args ...any) []any {
	httpArgs := []any{
		"method", request.Method,
		"url", request.URL,
	}

	allArgs := append(httpArgs, args...)

	return allArgs
}
