package usecases

import (
	"context"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g GuestUseCaseImpl) GetGuestsList(ctx context.Context) ([]models.Guest, error) {
	guests, err := g.guestRepository.GetGuestsList()
	if err != nil {
		return nil, err
	}
	return guests, nil
}
