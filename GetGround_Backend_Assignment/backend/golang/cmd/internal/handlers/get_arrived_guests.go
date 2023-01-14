package handlers

import (
	"encoding/json"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"net/http"
)

func (h *BaseHandler) GetArrivedGuests(w http.ResponseWriter, r *http.Request) {
	guests, err := h.groundService.GetArrivedGuests(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string][]models.ArrivedGuest{"guests": guests})
}
