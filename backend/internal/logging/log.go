package logging

import (
	"log/slog"
	"os"
)

func Init(config *Config) *slog.Logger {

	options := setUpHandlerOptions(config.Level)
	handler := setUpHandler(config.OutputType, options)
	logger := slog.New(handler)

	return logger
}

func setUpHandlerOptions(logLevel LogLevel) *slog.HandlerOptions {
	level := slog.LevelInfo

	switch logLevel {
	case LogLevelDebug:
		level = slog.LevelDebug
	case LogLevelInfo:
		level = slog.LevelInfo
	case LogLevelWarn:
		level = slog.LevelWarn
	case LogLevelError:
		level = slog.LevelError
	default:
		slog.Default().Warn("invalid log level value, using log level info", "logLevel", logLevel)
	}

	options := &slog.HandlerOptions{
		Level: level,
	}

	return options
}

func setUpHandler(outputType OutputType, options *slog.HandlerOptions) slog.Handler {
	var handler slog.Handler

	switch outputType {
	case OutputTypeJson:
		handler = slog.NewJSONHandler(os.Stdout, options)
	case OutputTypeText:
		handler = slog.NewTextHandler(os.Stdout, options)
	default:
		slog.Default().Warn("invalid output type value, defaulting to text output", "outputType", outputType)
		handler = slog.NewJSONHandler(os.Stdout, options)
	}

	return handler
}
