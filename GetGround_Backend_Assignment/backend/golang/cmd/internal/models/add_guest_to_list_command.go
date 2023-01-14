package models

type AddGuestToListCommand struct {
	Name               string `json:"name"`
	Table              int    `json:"table"`
	AccompanyingGuests int    `json:"accompanying_guests"`
}
