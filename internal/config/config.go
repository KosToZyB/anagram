package config

import "os"

const prefixEnv = "ANAGRAM_"

func GetEnv(name string) string {
	return os.Getenv(prefixEnv + name)
}

func SetEnv(key, value string) error {
	return os.Setenv(prefixEnv+key, value)
}

type Configer interface {
	Log() Logger
}

type Config struct {
	LogSetting      Log
}

func (c Config) Log() Logger {
	return c.LogSetting
}

func EnvConfig() Configer {
	return &Config{}
}
