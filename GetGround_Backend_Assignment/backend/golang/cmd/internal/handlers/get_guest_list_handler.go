package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *GuestHandler) GetGuestsList(w http.ResponseWriter, r *http.Request) {
	guests, err := h.groundService.GetGuestsList(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(guests)
}
