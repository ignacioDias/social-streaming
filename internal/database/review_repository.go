package database

import (
	"context"
	"errors"
	"socialstreaming/internal/domain"

	"github.com/jmoiron/sqlx"
)

type ReviewRepository struct {
	db *sqlx.DB
}

var ErrReviewNotFound = errors.New("Review Not Found")

func NewReviewRepository(db *sqlx.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}

func (reviewRepo *ReviewRepository) CreateReview(ctx context.Context, review *domain.Review) (int64, error) {
	query := `INSERT INTO reviews (user_id, score, review, target_id) VALUES ($1, $2, $3, $4) RETURNING review_id`
	var id int64
	err := reviewRepo.db.QueryRowContext(ctx, query, review.UserID, review.Score, review.Review, review.TargetID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (reviewRepo *ReviewRepository) DeleteReviewByID(ctx context.Context, reviewID int64) error {
	query := `DELETE FROM reviews WHERE review_id = $1`
	result, err := reviewRepo.db.ExecContext(ctx, query, reviewID)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrReviewNotFound
	}
	return nil
}

func (reviewRepo *ReviewRepository) GetAllReviewsPerUser(ctx context.Context, userID int64) ([]domain.Review, error) {
	query := `SELECT review_id, user_id, score, review, target_id FROM reviews WHERE user_id = $1`
	return reviewRepo.genericGetAllReviews(ctx, query, userID)
}

func (reviewRepo *ReviewRepository) GetAllReviewsPerMedia(ctx context.Context, mediaID int64) ([]domain.Review, error) {
	query := `SELECT review_id, user_id, score, review, target_id FROM reviews WHERE target_id = $1`
	return reviewRepo.genericGetAllReviews(ctx, query, mediaID)

}

func (reviewRepo *ReviewRepository) genericGetAllReviews(ctx context.Context, query string, idToSearch int64) ([]domain.Review, error) {
	var reviews []domain.Review
	if err := sqlx.SelectContext(ctx, reviewRepo.db, &reviews, query, idToSearch); err != nil {
		return nil, err
	}
	return reviews, nil
}
