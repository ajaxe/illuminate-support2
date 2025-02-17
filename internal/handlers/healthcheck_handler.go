package handlers

import (
	"net/http"
)

type HealthCheckHandler struct {
}
func (h *HealthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
