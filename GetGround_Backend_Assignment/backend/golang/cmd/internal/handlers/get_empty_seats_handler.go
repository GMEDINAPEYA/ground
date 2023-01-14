package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *BaseHandler) GetEmptySeats(w http.ResponseWriter, r *http.Request) {
	emptySeats, err := h.groundService.GetEmptySeats(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"seats_empty": emptySeats})
}
