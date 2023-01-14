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
