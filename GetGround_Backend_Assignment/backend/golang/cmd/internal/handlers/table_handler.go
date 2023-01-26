package handlers

import "github.com/getground/tech-tasks/backend/cmd/internal/usecases"

type TableHandler struct {
	tableUseCase usecases.TableUseCase
}

func NewTableHandler(tableUseCase usecases.TableUseCase) *TableHandler {
	return &TableHandler{
		tableUseCase: tableUseCase,
	}
}
