package database

import (
	"context"
	"errors"
	"socialstreaming/internal/domain"

	"github.com/jmoiron/sqlx"
)

var ErrFollowNotFound error = errors.New("Follow Not Found")

type FollowRepository struct {
	db *sqlx.DB
}

func NewFollowRepository(db *sqlx.DB) *FollowRepository {
	return &FollowRepository{db: db}
}

func (fwRepo *FollowRepository) CreateFollow(ctx context.Context, followerID int64, followedID int64) (int64, error) {
	query := `INSERT INTO follows (follower_id, followed_id) VALUES ($1, $2) RETURNING follow_id`
	var id int64
	err := fwRepo.db.QueryRowContext(ctx, query, followerID, followedID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (fwRepo *FollowRepository) RemoveFollow(ctx context.Context, followerID int64, followedID int64) error {
	query := `DELETE FROM follows WHERE follower_id = $1 AND followed_id = $2`
	result, err := fwRepo.db.ExecContext(ctx, query, followerID, followedID)
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
func (fwRepo *FollowRepository) GetAllFollowsPerUser(ctx context.Context, userID int64) ([]domain.Follow, error) {
	query := `SELECT follow_id, follower_id, followed_id FROM follows WHERE follower_id = $1`
	return fwRepo.genericGetAllPerUser(ctx, userID, query)
}
func (fwRepo *FollowRepository) GetAllFollowersPerUser(ctx context.Context, userID int64) ([]domain.Follow, error) {
	query := `SELECT follow_id, follower_id, followed_id FROM follows WHERE followed_id = $1`
	return fwRepo.genericGetAllPerUser(ctx, userID, query)
}

func (fwRepo *FollowRepository) genericGetAllPerUser(ctx context.Context, userID int64, query string) ([]domain.Follow, error) {
	var follows []domain.Follow
	if err := sqlx.SelectContext(ctx, fwRepo.db, &follows, query, userID); err != nil {
		return nil, err
	}
	return follows, nil
}

func (fwRepo *FollowRepository) CheckIfUserFollowsAnother(ctx context.Context, followerID int64, followedID int64) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = $1 AND followed_id = $2)`
	var exists bool
	err := fwRepo.db.QueryRowContext(ctx, query, followerID, followedID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
