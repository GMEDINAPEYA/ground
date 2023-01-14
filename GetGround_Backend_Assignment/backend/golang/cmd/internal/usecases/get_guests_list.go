package usecases

import (
	"context"
	"fmt"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func (g GetGroundServiceImpl) GetGuestsList(ctx context.Context) ([]models.Guest, error) {
	guests, err := g.guestRepository.GetGuestsList()
	if err != nil {
		fmt.Println("Error getting guest")
	}
	return guests, nil
}
