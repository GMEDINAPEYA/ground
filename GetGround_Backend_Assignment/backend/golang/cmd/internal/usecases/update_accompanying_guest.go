package usecases

import (
	"context"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g GetGroundServiceImpl) UpdateAccompanyingGuests(ctx context.Context, name string, req models.UpdateAccompanyingGuestsCommand) error {
	err := g.guestRepository.UpdateAccompanyingGuests(name, req.AccompanyingGuests)
	if err != nil {
		return err
	}
	return nil

}
