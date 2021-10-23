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

// HTTP controller
func (health *Health) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if(r.Method == http.MethodGet) {
		health.logger.Println("Handle Health GET")
		// d, err := json.Marshal("OK");
		// if err != nil {
		// 	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		// }
		// rw.Write(d)
		rw.Write([]byte("OK"))
	}
	health.logger.Println("OK")
}