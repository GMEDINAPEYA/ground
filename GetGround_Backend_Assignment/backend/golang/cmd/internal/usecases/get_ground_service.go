package usecases

import (
	"context"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/getground/tech-tasks/backend/cmd/internal/repositories"
)

type GetGroundService interface {
	CreateTable(ctx context.Context, cmd models.CreateTableCommand) (*models.Table, error)
	AddGuestToList(ctx context.Context, cmd models.AddGuestToListCommand) error
	GetGuestsList(ctx context.Context) ([]models.Guest, error)
	UpdateAccompanyingGuests(ctx context.Context, name string, req models.UpdateAccompanyingGuestsCommand) error
	DeleteGuest(ctx context.Context, name string) error
	GetArrivedGuests(ctx context.Context) ([]models.ArrivedGuest, error)
	GetEmptySeats(ctx context.Context) (int, error)
}

type GetGroundServiceImpl struct {
	tableRepository repositories.TableRepo
	guestRepository repositories.GuestRepo
}

func NewGetGroundServiceService(
	tableRepo repositories.TableRepo,
	guestRepo repositories.GuestRepo,
) GetGroundService {
	return &GetGroundServiceImpl{
		tableRepository: tableRepo,
		guestRepository: guestRepo,
	}
}
