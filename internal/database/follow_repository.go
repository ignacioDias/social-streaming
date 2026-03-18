package database

import (
	"context"
	"errors"
	"socialstreaming/internal/domain"

	"github.com/jmoiron/sqlx"
)

var ErrFollowNotFound = errors.New("Follow Not Found")

type FollowRepository struct {
	db *sqlx.DB
}

func NewFollowRepository(db *sqlx.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

func (fr *FollowRepository) CreateFollow(ctx context.Context, follow *domain.Follow) (int64, error) {
	query := `INSERT INTO follows (follower_id, followed_id) VALUES ($1, $2) RETURNING follow_id`
	var id int64
	err := fr.db.QueryRowContext(ctx, query, follow.FollowerID, follow.FollowedID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (fr *FollowRepository) RemoveFollow(ctx context.Context, follow *domain.Follow) error {
	query := `DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2`
	result, err := fr.db.ExecContext(ctx, query, follow.FollowerID, follow.FollowedID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrFollowNotFound
	}
	return nil
}
func (fr *FollowRepository) GetAllFollowsPerUser(ctx context.Context, userID int64, limit, offset int) ([]domain.Follow, error) {
	query := `SELECT follow_id, follower_id, followed_id FROM follows WHERE follower_id = $1 ORDER BY created_at LIMIT $2 OFFSET $3`
	return fr.genericGetAllPerUser(ctx, userID, query, limit, offset)
}
func (fr *FollowRepository) GetAllFollowersPerUser(ctx context.Context, userID int64, limit, offset int) ([]domain.Follow, error) {
	query := `SELECT follow_id, follower_id, followed_id FROM follows WHERE followed_id = $1 ORDER BY created_at LIMIT $2 OFFSET $3`
	return fr.genericGetAllPerUser(ctx, userID, query, limit, offset)
}

func (fr *FollowRepository) genericGetAllPerUser(ctx context.Context, userID int64, query string, limit, offset int) ([]domain.Follow, error) {
	var follows []domain.Follow
	if err := sqlx.SelectContext(ctx, fr.db, &follows, query, userID, limit, offset); err != nil {
		return nil, err
	}
	return follows, nil
}

func (fr *FollowRepository) CheckIfUserFollowsAnother(ctx context.Context, follow *domain.Follow) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2)`
	var exists bool
	err := fr.db.QueryRowContext(ctx, query, follow.FollowerID, follow.FollowedID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
