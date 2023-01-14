package usecases

import (
	"context"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g GetGroundServiceImpl) GetArrivedGuests(ctx context.Context) ([]models.ArrivedGuest, error) {
	guests, err := g.guestRepository.GetArrivedGuest()
	if err != nil {
		return nil, err
	}

	return guests, nil
}
