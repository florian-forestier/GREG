package controller

import (
	"encoding/json"
	"github.com/Artheriom/GREG/internal/service"
	"github.com/Artheriom/GREG/internal/structures"
	"net/http"
)

func AddMetric(w http.ResponseWriter, r *http.Request) {

	var item structures.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
	}

	service.Append(item)

}
