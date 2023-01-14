package handlers

import (
	"encoding/json"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"net/http"
)

func (h *BaseHandler) GetGuestsList(w http.ResponseWriter, r *http.Request) {
	guests, err := h.groundService.GetGuestsList(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string][]models.Guest{"guests": guests})
}
