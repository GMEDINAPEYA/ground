package mocks

import (
	"context"

	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/stretchr/testify/mock"
)

type GetGroundService struct {
	mock.Mock
}

func (m *GetGroundService) CreateTable(ctx context.Context, cmd models.CreateTableCommand) (*models.Table, error) {
	args := m.Called(ctx, cmd)
	return args.Get(0).(*models.Table), args.Error(1)
}

func (m *GetGroundService) AddGuestToList(ctx context.Context, cmd models.AddGuestToListCommand) error {
	args := m.Called(ctx, cmd)
	return args.Error(0)
}

func (m *GetGroundService) GetGuestsList(ctx context.Context) ([]models.Guest, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Guest), args.Error(1)
}

func (m *GetGroundService) UpdateAccompanyingGuests(ctx context.Context, name string, req models.UpdateAccompanyingGuestsCommand) error {
	args := m.Called(ctx, name, req)
	return args.Error(0)
}

func (m *GetGroundService) DeleteGuest(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func (m *GetGroundService) GetArrivedGuests(ctx context.Context) ([]models.ArrivedGuest, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.ArrivedGuest), args.Error(1)
}

func (m *GetGroundService) GetEmptySeats(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

type GetGroundServiceImpl struct {
	mock.Mock
}

func (m *GetGroundServiceImpl) CreateTable(ctx context.Context, cmd models.CreateTableCommand) (*models.Table, error) {
	args := m.Called(ctx, cmd)
	return args.Get(0).(*models.Table), args.Error(1)
}

func (m *GetGroundServiceImpl) AddGuestToList(ctx context.Context, cmd models.AddGuestToListCommand) error {
	args := m.Called(ctx, cmd)
	return args.Error(0)
}

func (m *GetGroundServiceImpl) GetGuestsList(ctx context.Context) ([]models.Guest, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Guest), args.Error(1)
}
