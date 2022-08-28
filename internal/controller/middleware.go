package controller

import (
	"github.com/Artheriom/GREG/internal/helpers"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Vary", "Origin")
		w.Header().Add("Vary", "Access-Control-Request-Method")
		w.Header().Add("Vary", "Access-Control-Request-Headers")
		w.Header().Add("Access-Control-Expose-Headers", "Authorization")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,OPTIONS")

		if r.Method == "OPTIONS" {
			return
		}

		if helpers.Bearers[r.Method] != "" && helpers.Bearers[r.Method] != r.Header.Get("Authorization") {
			w.WriteHeader(403)
			return
		}

		next.ServeHTTP(w, r)

	})
}
