package usecases

import (
	"context"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/getground/tech-tasks/backend/cmd/internal/repositories"
)

type GuestUseCase interface {
	AddGuestToList(ctx context.Context, cmd models.AddGuestToListCommand) error
	GetGuestsList(ctx context.Context) ([]models.Guest, error)
	UpdateAccompanyingGuests(ctx context.Context, name string, req models.UpdateAccompanyingGuestsCommand) error
	DeleteGuest(ctx context.Context, name string) error
	GetArrivedGuests(ctx context.Context) ([]models.ArrivedGuest, error)
}

type GuestUseCaseImpl struct {
	tableRepository repositories.TableRepo
	guestRepository repositories.GuestRepo
}

func NewGuestUseCase(
	guestRepo repositories.GuestRepo,
	tableRepo repositories.TableRepo,
) GuestUseCase {
	return &GuestUseCaseImpl{
		guestRepository: guestRepo,
		tableRepository: tableRepo,
	}
}
