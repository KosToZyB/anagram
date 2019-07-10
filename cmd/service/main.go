package main

import (
	"github.com/kostozyb/anagram/internal/config"
	"github.com/kostozyb/anagram/internal/server"
	"github.com/kostozyb/anagram/pkg/logger"

	_ "github.com/lib/pq"

	lg "log"
)

func main() {
	env := config.EnvConfig()

	logLevel := env.Log().GetLevel()
	log, err := logger.GetLogger(logLevel)
	if err != nil {
		lg.Panicf("error initialize custom logger: %s\n", err)
	}

	s := server.NewServer(env, log)
	s.Run()
}
