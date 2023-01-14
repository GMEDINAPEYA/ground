package usecases

import (
	"context"
	"errors"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g GetGroundServiceImpl) UpdateAccompanyingGuests(ctx context.Context, name string, req models.UpdateAccompanyingGuestsCommand) error {
	// Get guest info
	guest, err := g.guestRepository.GetGuest(name)
	if err != nil {
		return err
	}

	// Get table capacity
	table, err := g.tableRepository.GetTableInfo(guest.Table)
	if err != nil {
		return err
	}

	// Recalculate table capacity
	originalTotalGuests := guest.AccompanyingGuests + 1
	originalTableCapacity := table.Capacity + originalTotalGuests
	newTotalGuests := req.AccompanyingGuests + 1
	newTableCapacity := originalTableCapacity - newTotalGuests

	if newTableCapacity < 0 {
		msg := fmt.Sprintf("guests amount: %v exceed table capacity: %v", req.AccompanyingGuests, originalTableCapacity)
		return errors.New(msg)
	}

	err = g.guestRepository.UpdateAccompanyingGuests(name, req.AccompanyingGuests)
	if err != nil {
		return err
	}

	err = g.tableRepository.UpdateTableCapacity(guest.Table, newTableCapacity)
	if err != nil {
		return err
	}

	return nil
}
