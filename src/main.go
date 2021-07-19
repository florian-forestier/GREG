package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"controller"
)

func main() {
		r := mux.NewRouter()
		
		api := r.PathPrefix("").Subrouter()
		api.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {controller.GetMetrics(w, r)}).Methods(http.MethodGet, http.MethodOptions)
		api.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {controller.AddMetric(w, r)}).Methods(http.MethodPost, http.MethodOptions)
		
		fmt.Println(http.ListenAndServe(":9666", r))

}