package repositories

import (
	"database/sql"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

type TableRepo struct {
	db *sql.DB
}

func NewTableRepo(db *sql.DB) *TableRepo {
	return &TableRepo{
		db: db,
	}
}

type TableRepository interface {
	Save(table *models.Table) (*models.Table, error)
	GetTableInfo(id int) (*models.Table, error)
	UpdateTableCapacity(id int, amountOfGuests int) error
	GetEmptySeats() (int, error)
}

func (r *TableRepo) GetEmptySeats() (int, error) {
	query := "SELECT * FROM tables"
	rows, err := r.db.Query(query)

	if err != nil {
		return 0, err
	}

	var emptySeats int

	for rows.Next() {
		var table models.Table
		err = rows.Scan(&table.Id, &table.Capacity)
		if err != nil {
			return 0, err
		}
		emptySeats += table.Capacity
	}

	return emptySeats, nil
}

func (r *TableRepo) UpdateTableCapacity(id int, newCapacity int) error {
	query := "UPDATE tables SET capacity = ? where id = ?"
	_, err := r.db.Query(query, newCapacity, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *TableRepo) GetTableInfo(id int) (*models.Table, error) {
	query := "SELECT * FROM tables WHERE id = ?"

	var table models.Table
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&table.Id, &table.Capacity); err != nil {
		return nil, err
	}

	return &table, nil
}

func (r *TableRepo) Save(table *models.Table) (*models.Table, error) {
	query := "INSERT INTO tables(capacity) VALUES (?)"
	row, err := r.db.Exec(query, table.Capacity)
	id, err := row.LastInsertId()

	if err != nil {
		return nil, err
	}

	result := models.Table{
		Id:       int(id),
		Capacity: table.Capacity,
	}

	return &result, err
}
