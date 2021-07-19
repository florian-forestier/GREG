package controller

import (
	"net/http"
	"service"
)

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Vary", "Origin")
	w.Header().Add("Vary", "Access-Control-Request-Method")
	w.Header().Add("Vary", "Access-Control-Request-Headers")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")
	w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS")

	if r.Method == "OPTIONS" {
		return
	}
	
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	
	w.Write([]byte(service.ParseCache()))
}