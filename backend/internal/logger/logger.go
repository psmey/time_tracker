package logger

import (
	"log/slog"
	"os"
)

type Config struct {
	OutputType OutputType `yaml:"outputType"`
	Level      LogLevel   `yaml:"logLevel"`
}

type OutputType string

const (
	OutputTypeJson OutputType = "json"
	OutputTypeText OutputType = "text"
)

type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

var logger *slog.Logger

func Init(config *Config) {
	options := setUpHandlerOptions(config.Level)
	handler := setUpHandler(config.OutputType, options)
	logger = slog.New(handler)
}

func LogDebug(message string, args ...any) {
	ensureLogger().Debug(message, args...)
}

func LogInfo(message string, args ...any) {
	ensureLogger().Info(message, args...)
}

func LogWarn(message string, args ...any) {
	ensureLogger().Warn(message, args...)
}

func LogError(message string, err error, args ...any) {
	allArgs := append([]any{"error", err}, args...)
	ensureLogger().Error(message, allArgs...)
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

func ensureLogger() *slog.Logger {
	if logger == nil {
		return slog.Default()
	}
	return logger
}
