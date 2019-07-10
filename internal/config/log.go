package config

const (
	logLevel = "LOG_LEVEL"
)

type Logger interface {
	GetLevel() string
}

type Log struct {
}

func (l Log) GetLevel() string {
	return GetEnv(logLevel)
}
