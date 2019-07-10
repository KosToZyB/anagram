package middleware

import (
	"github.com/kostozyb/anagram/pkg/logger"
)

type Middleware struct {
	log logger.Logger
}

func NewMiddleware(log logger.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}
