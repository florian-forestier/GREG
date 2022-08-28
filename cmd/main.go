package main

import (
	"fmt"
	"github.com/Artheriom/GREG/internal/controller"
	"github.com/Artheriom/GREG/internal/helpers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	api := r.PathPrefix(helpers.PathPrefix).Subrouter()
	api.Use(controller.Middleware)

	api.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { controller.GetMetrics(w, r) }).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) { controller.AddMetric(w, r) }).Methods(http.MethodPost, http.MethodOptions)

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", helpers.Port), r))

}
