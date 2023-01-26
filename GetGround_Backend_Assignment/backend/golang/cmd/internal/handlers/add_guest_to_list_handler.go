package handlers

import (
	"encoding/json"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *GuestHandler) AddGuestToGuestList(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	name := param["name"]

	var req models.AddGuestToListCommand

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Table <= 0 {
		http.Error(w, "invalid table ID. Must be greater than 0", http.StatusBadRequest)
		return
	}

	if req.AccompanyingGuests < 0 {
		http.Error(w, "accompanying guests can not be smaller than 0", http.StatusBadRequest)
		return
	}

	req.Name = name

	if err = h.groundService.AddGuestToList(r.Context(), req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"name": name})
}
