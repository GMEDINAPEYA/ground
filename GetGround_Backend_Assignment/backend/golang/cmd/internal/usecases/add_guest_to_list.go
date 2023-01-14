package usecases

import (
	"context"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g GetGroundServiceImpl) AddGuestToList(ctx context.Context, cmd models.AddGuestToListCommand) error {
	// Validate capacity
	guest := &models.Guest{
		Name:               cmd.Name,
		Table:              cmd.Table,
		AccompanyingGuests: cmd.AccompanyingGuests,
	}

	err := g.guestRepository.Save(guest)
	if err != nil {
		fmt.Println("Error adding guest")
	}
	return nil
}
