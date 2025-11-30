package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/psmey/time_tracker/internal/features/calendar/domain"
)

type RepositoryPort interface {
	DeleteCalendar(id uuid.UUID) error
}

type PostgresRepositoryAdapter struct {
	db *sql.DB
}

var _ RepositoryPort = &PostgresRepositoryAdapter{}

func NewPostgresRepository(db *sql.DB) *PostgresRepositoryAdapter {
	return &PostgresRepositoryAdapter{db: db}
}

func (repository *PostgresRepositoryAdapter) DeleteCalendar(id uuid.UUID) error {
	result, err := repository.db.ExecContext(
		context.Background(),
		`DELETE FROM calendars WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return domain.ErrCalendarNotFound
	}

	return nil
}
