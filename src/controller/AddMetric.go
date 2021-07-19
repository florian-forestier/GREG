package controller

import (
	"net/http"
	"service"
	"encoding/json"
)

func AddMetric(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Vary", "Origin")
	w.Header().Add("Vary", "Access-Control-Request-Method")
	w.Header().Add("Vary", "Access-Control-Request-Headers")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")
	w.Header().Add("Access-Control-Allow-Methods", "POST,OPTIONS")

	if r.Method == "OPTIONS" {
		return
	}
	
	var item service.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}
	
	
	service.Append(item)


}