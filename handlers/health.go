package handlers

import (
	"log"
	"net/http"
)

type Health struct {
	logger *log.Logger
}

// Dependency injection
func NewHealth(logger *log.Logger) *Health {
	return &Health{logger}
}

func (health *Health) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	health.logger.Println("OK")
}