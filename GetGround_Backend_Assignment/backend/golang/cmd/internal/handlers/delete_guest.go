package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (h *GuestHandler) DeleteGuest(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	name := param["name"]

	if name == "" {
		http.Error(w, "missing user name. Name can not be empty", http.StatusBadRequest)
		return
	}

	err := h.groundService.DeleteGuest(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
