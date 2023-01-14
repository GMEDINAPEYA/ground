package models

type Guest struct {
	Name               string `json:"name"`
	Table              int    `json:"table"`
	AccompanyingGuests int    `json:"accompanying_guests"`
}

type ArrivedGuest struct {
	Name               string `json:"name"`
	AccompanyingGuests int    `json:"accompanying_guests"`
	TimeArrived        string `json:"time_arrived"`
}
