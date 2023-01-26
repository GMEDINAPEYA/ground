package handlers

import "github.com/getground/tech-tasks/backend/cmd/internal/usecases"

type GuestHandler struct {
	groundService usecases.GuestUseCase
}

func NewGuestHandler(groundService usecases.GuestUseCase) *GuestHandler {
	return &GuestHandler{
		groundService: groundService,
	}
}
