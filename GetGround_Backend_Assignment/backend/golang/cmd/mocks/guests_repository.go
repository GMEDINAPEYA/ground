package mocks

import (
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"github.com/stretchr/testify/mock"
)

// GuestRepositoryMock is a mock implementation of the GuestRepository interface
type GuestRepositoryMock struct {
	mock.Mock
}

func (m *GuestRepositoryMock) Save(guest *models.Guest) error {
	args := m.Called(guest)
	return args.Error(0)
}

func (m *GuestRepositoryMock) GetGuestsList() ([]models.Guest, error) {
	args := m.Called()
	return args.Get(0).([]models.Guest), args.Error(1)
}

func (m *GuestRepositoryMock) UpdateAccompanyingGuests(name string, amount int) error {
	args := m.Called(name, amount)
	return args.Error(0)
}

func (m *GuestRepositoryMock) DeleteGuest(name string) error {
	args := m.Called(name)
	return args.Error(0)
}
