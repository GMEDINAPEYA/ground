package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g GuestUseCaseImpl) AddGuestToList(ctx context.Context, cmd models.AddGuestToListCommand) error {
	// Validate table capacity
	tableInfo, err := g.tableRepository.GetTableInfo(cmd.Table)

	if err != nil {
		return err
	}

	totalGuests := cmd.AccompanyingGuests + 1

	if cmd.AccompanyingGuests+1 > tableInfo.Capacity {
		msg := fmt.Sprintf("guests amount: %v exceed table capacity: %v", totalGuests, tableInfo.Capacity)
		return errors.New(msg)
	}

	guest := &models.Guest{
		Name:               cmd.Name,
		Table:              cmd.Table,
		AccompanyingGuests: cmd.AccompanyingGuests,
	}

	err = g.guestRepository.Save(guest)
	if err != nil {
		return err
	}

	newCapacity := tableInfo.Capacity - totalGuests

	// Update table capacity
	err = g.tableRepository.UpdateTableCapacity(cmd.Table, newCapacity)
	if err != nil {
		return err
	}

	return nil
}
