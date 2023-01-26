package usecases

import "context"

func (g GuestUseCaseImpl) DeleteGuest(ctx context.Context, name string) error {
	// Get guest first in order to calculate total amount of guests
	guest, err := g.guestRepository.GetGuest(name)
	if err != nil {
		return err
	}

	totalGuests := guest.AccompanyingGuests + 1

	err = g.guestRepository.DeleteGuest(name)
	if err != nil {
		return err
	}

	// Update table capacity
	err = g.tableRepository.UpdateTableCapacity(guest.Table, totalGuests)
	if err != nil {
		return err
	}

	return nil
}
