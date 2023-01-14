package handlers

import "github.com/getground/tech-tasks/backend/cmd/internal/usecases"

type BaseHandler struct {
	groundService usecases.GetGroundService
}

func NewBaseHandler(groundService usecases.GetGroundService) *BaseHandler {
	return &BaseHandler{
		groundService: groundService,
	}
}
