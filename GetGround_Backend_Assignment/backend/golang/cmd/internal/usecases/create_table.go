package usecases

import (
	"context"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g TableUseCaseImpl) CreateTable(ctx context.Context, cmd models.CreateTableCommand) (*models.Table, error) {
	table := &models.Table{
		Capacity: cmd.Capacity,
	}

	t, err := g.tableRepository.Save(table)
	if err != nil {
		return nil, err
	}

	return t, nil
}
