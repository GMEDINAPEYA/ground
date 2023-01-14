package repositories

import (
	"database/sql"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
	"time"
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
	GetGuest(name string) (*models.Guest, error)
	GetArrivedGuests() ([]models.ArrivedGuest, error)
}

func (r *GuestRepo) GetGuest(name string) (*models.Guest, error) {
	query := "SELECT guest_name, table_id, accompanying_guests FROM guests where guest_name = ?"

	var guest models.Guest
	row := r.db.QueryRow(query, name)

	if err := row.Scan(&guest.Name, &guest.Table, &guest.AccompanyingGuests); err != nil {
		return nil, err
	}

	return &guest, nil
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
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	query := "INSERT INTO guests(guest_name, table_id, accompanying_guests, time_arrived) VALUES (?, ?, ?, ?)"
	_, err := r.db.Query(query, guest.Name, guest.Table, guest.AccompanyingGuests, timestamp)

	if err != nil {
		return err
	}

	return nil
}

func (r *GuestRepo) GetArrivedGuest() ([]models.ArrivedGuest, error) {
	query := "SELECT guest_name, accompanying_guests, time_arrived FROM guests"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	var arrivedGuests []models.ArrivedGuest

	for rows.Next() {
		var guest models.ArrivedGuest
		err = rows.Scan(&guest.Name, &guest.AccompanyingGuests, &guest.TimeArrived)
		if err != nil {
			return nil, err
		}
		arrivedGuests = append(arrivedGuests, guest)
	}

	return arrivedGuests, nil
}

func (r *GuestRepo) GetGuestsList() ([]models.Guest, error) {
	query := "SELECT guest_name, table_id, accompanying_guests FROM guests"
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
