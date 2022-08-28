package controller

import (
	"github.com/Artheriom/GREG/internal/service"
	"net/http"
)

func GetMetrics(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "text/plain; version=0.0.4")

	_, _ = w.Write([]byte(service.ParseCache()))
}
