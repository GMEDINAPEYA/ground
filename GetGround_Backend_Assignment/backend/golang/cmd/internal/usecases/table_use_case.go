package usecases

import (
	"context"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/getground/tech-tasks/backend/cmd/internal/repositories"
)

type TableUseCase interface {
	CreateTable(ctx context.Context, cmd models.CreateTableCommand) (*models.Table, error)
	GetEmptySeats(ctx context.Context) (int, error)
}

type TableUseCaseImpl struct {
	tableRepository repositories.TableRepo
}

func NewTableUseCase(
	tableRepo repositories.TableRepo,
) TableUseCase {
	return &TableUseCaseImpl{
		tableRepository: tableRepo,
	}
}
