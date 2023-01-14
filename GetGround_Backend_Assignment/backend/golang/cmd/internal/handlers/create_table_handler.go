package handlers

import (
	"encoding/json"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"net/http"
)

func (h *BaseHandler) CreateTable(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTableCommand

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Capacity <= 0 {
		http.Error(w, "invalid capacity provided. Must be greater than 0", http.StatusBadRequest)
		return
	}

	t, err := h.groundService.CreateTable(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(t)
}
