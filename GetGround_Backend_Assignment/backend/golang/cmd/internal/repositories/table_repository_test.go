package repositories

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/getground/tech-tasks/backend/cmd/internal/models"
)

func TestGetEmptySeats(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &TableRepo{db: db}

	// test empty seats
	mock.ExpectQuery("SELECT \\* FROM tables").WillReturnRows(sqlmock.NewRows([]string{"id", "capacity"}).AddRow(1, 5).AddRow(2, 3))
	emptySeats, err := r.GetEmptySeats()
	if err != nil {
		t.Errorf("error was not expected while getting empty seats: %s", err)
	}
	if emptySeats != 8 {
		t.Errorf("empty seats count not matched, expected: 8, got: %d", emptySeats)
	}

	// test error scenario
	mock.ExpectQuery("SELECT \\* FROM tables").WillReturnError(sql.ErrNoRows)
	_, err = r.GetEmptySeats()
	if err == nil {
		t.Error("error was expected while getting empty seats, but got nil")
	}
}

func TestUpdateTableCapacity(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &TableRepo{db: db}

	// test successful update
	mock.ExpectQuery("UPDATE tables SET capacity = \\? where id = \\?").WithArgs(10, 1).WillReturnRows(sqlmock.NewRows([]string{"id", "capacity"}).AddRow(1, 5))
	err = r.UpdateTableCapacity(1, 10)
	if err != nil {
		t.Errorf("error was not expected while updating table capacity: %s", err)
	}

	// test error scenario
	mock.ExpectExec("UPDATE tables SET capacity = \\? where id = \\?").WithArgs(10, 1).WillReturnError(errors.New("forced for test"))
	err = r.UpdateTableCapacity(2, 10)
	if err == nil {
		t.Error("forced for test")
	}
}

func TestGetTableInfo(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &TableRepo{db: db}

	// test getting table info
	mock.ExpectQuery("SELECT \\* FROM tables WHERE id = \\?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "capacity"}).AddRow(1, 5))
	table, err := r.GetTableInfo(1)
	if err != nil {
		t.Errorf("error was not expected while getting table info: %s", err)
	}
	if table.Id != 1 || table.Capacity != 5 {
		t.Errorf("unexpected table info, expected (id: 1, capacity: 5), got: %v", table)
	}

	// test error scenario
	mock.ExpectQuery("SELECT \\* FROM tables WHERE id = \\?").WithArgs(2).WillReturnError(sql.ErrNoRows)
	_, err = r.GetTableInfo(2)
	if err == nil {
		t.Error("error was expected while getting table info, but got nil")
	}
}

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &TableRepo{db: db}

	// test saving table
	mock.ExpectExec("INSERT INTO tables\\(capacity\\) VALUES \\(\\?").WithArgs(5).WillReturnResult(sqlmock.NewResult(1, 1))
	table := &models.Table{Capacity: 5}
	savedTable, err := r.Save(table)
	if err != nil {
		t.Errorf("error was not expected while saving table: %s", err)
	}
	if savedTable.Id != 1 || savedTable.Capacity != 5 {
		t.Errorf("unexpected saved table, expected (id: 1, capacity: 5), got: %v", savedTable)
	}
}
