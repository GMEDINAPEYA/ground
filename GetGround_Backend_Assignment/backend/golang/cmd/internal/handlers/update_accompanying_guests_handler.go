package handlers

import (
	"encoding/json"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *BaseHandler) UpdateAccompanyingGuests(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	name := param["name"]
	var req models.UpdateAccompanyingGuestsCommand

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if name == "" {
		http.Error(w, "missing user name. Name can not be empty", http.StatusBadRequest)
		return
	}

	if req.AccompanyingGuests <= 0 {
		http.Error(w, "you should add at least one more guest", http.StatusBadRequest)
		return
	}

	err = h.groundService.UpdateAccompanyingGuests(r.Context(), name, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"name": name})
}
