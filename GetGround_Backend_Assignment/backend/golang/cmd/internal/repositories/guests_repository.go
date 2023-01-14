package repositories

import (
	"database/sql"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

type GuestRepo struct {
	db *sql.DB
}

func NewGuestRepo(db *sql.DB) *GuestRepo {
	return &GuestRepo{
		db: db,
	}
}

type GuestRepository interface {
	Save(guest *models.Guest) error
	GetGuestsList() ([]models.Guest, error)
	UpdateAccompanyingGuests(name string, amount int) error
	DeleteGuest(name string) error
}

func (r *GuestRepo) DeleteGuest(name string) error {
	query := "DELETE FROM guests WHERE guest_name = ?"
	_, err := r.db.Query(query, name)

	if err != nil {
		return err
	}

	return nil
}

func (r *GuestRepo) UpdateAccompanyingGuests(name string, amount int) error {
	query := "UPDATE guests SET accompanying_guests = ? where guest_name = ?"
	_, err := r.db.Query(query, amount, name)

	if err != nil {
		return err
	}

	return nil
}

func (r *GuestRepo) Save(guest *models.Guest) error {
	query := "INSERT INTO guests(guest_name, table_id, accompanying_guests) VALUES (?, ?, ?)"
	_, err := r.db.Query(query, guest.Name, guest.Table, guest.AccompanyingGuests)

	if err != nil {
		return err
	}

	return nil
}

func (r *GuestRepo) GetGuestsList() ([]models.Guest, error) {
	query := "SELECT * FROM guests"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var guests []models.Guest

	for rows.Next() {
		var guest models.Guest
		err = rows.Scan(&guest.Name, &guest.Table, &guest.AccompanyingGuests)
		if err != nil {
			return nil, err
		}
		guests = append(guests, guest)
	}

	return guests, nil
}
