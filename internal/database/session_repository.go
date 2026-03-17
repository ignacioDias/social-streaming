package database

import (
	"context"
	"database/sql"
	"errors"
	"socialstreaming/internal/domain"

	"github.com/jmoiron/sqlx"
)

type SessionRepository struct {
	db *sqlx.DB
}

var ErrSessionNotFound = errors.New("Session not found")

func (sessRepo *SessionRepository) CreateSession(ctx context.Context, session *domain.Session) error {
	query := `
	INSERT INTO sessions (id, user_id, expires_at)
	VALUES (:id, :user_id, :expires_at)
	RETURNING created_at
	`
	stmt, err := sessRepo.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return err
	}
	return stmt.GetContext(ctx, session, session)
}

func (sessRepo *SessionRepository) FindSessionByID(ctx context.Context, id string) (*domain.Session, error) {
	var session domain.Session
	query := "SELECT id, user_id, created_at, expires_at FROM sessions WHERE id = $1 AND expires_at > (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')"
	if err := sessRepo.db.GetContext(ctx, &session, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrSessionNotFound
		}
		return nil, err
	}
	return &session, nil
}

func (sessRepo *SessionRepository) DeleteSessionByID(ctx context.Context, id string) error {
	query := "DELETE FROM sessions WHERE id = $1"
	result, err := sessRepo.db.ExecContext(ctx, query, id)
	ret := checkQueryResult(result, err)
	return ret
}

func (sessRepo *SessionRepository) DeleteSessionsByUserID(ctx context.Context, userID int64) error {
	query := "DELETE FROM sessions WHERE user_id = $1"
	result, err := sessRepo.db.ExecContext(ctx, query, userID)
	ret := checkQueryResult(result, err)
	return ret
}

func checkQueryResult(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrSessionNotFound
	}
	return nil
}
