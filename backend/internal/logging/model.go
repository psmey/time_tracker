package logging

type Config struct {
	OutputType OutputType `yaml:"outputType"`
	Level      LogLevel   `yaml:"level"`
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
