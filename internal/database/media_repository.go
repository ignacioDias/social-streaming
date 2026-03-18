package database

import (
	"context"
	"socialstreaming/internal/domain"

	"github.com/jmoiron/sqlx"
)

type MediaRepository struct {
	db *sqlx.DB
}

func NewMediaRepository(db *sqlx.DB) *MediaRepository {
	return &MediaRepository{db: db}
}

func (mr *MediaRepository) CreateMovie(ctx context.Context, movie *domain.Movie) error {
	query := `INSERT INTO media (title, description, total_ratings, sum_ratings, tags) VALUES ($1, $2, 0, 0, $3) returning media_id`
	var mediaID int64
	if err := mr.db.QueryRowxContext(ctx, query, movie.Title, movie.Description, movie.Tags).Scan(&mediaID); err != nil {
		return err
	}
	movie.MediaID = mediaID
	queryMovie := `INSERT INTO movies (media_id, video_url) VALUES ($1, $2)`
	if _, err := mr.db.ExecContext(ctx, queryMovie, movie.MediaID, movie.VideoURL); err != nil {
		return err
	}
	return nil
}
func (mr *MediaRepository) CreateSeries(ctx context.Context, series *domain.Series) error {
	query := `INSERT INTO media (title, description, total_ratings, sum_ratings, tags) VALUES ($1, $2, 0, 0, $3) returning media_id`
	var mediaID int64
	if err := mr.db.QueryRowxContext(ctx, query, series.Title, series.Description, series.Tags).Scan(&mediaID); err != nil {
		return err
	}
	series.MediaID = mediaID
	queryMovie := `INSERT INTO series (media_id) VALUES ($1)`
	if _, err := mr.db.ExecContext(ctx, queryMovie, mediaID); err != nil {
		return err
	}
	return nil
}

func (mr *MediaRepository) CreateSeason(ctx context.Context, season *domain.Season) error {
	query := `INSERT INTO seasons (title, number, series_id) VALUES ($1, $2, $3) returning season_id`
	var id int64
	if err := mr.db.QueryRowxContext(ctx, query, season.Title, season.Number, season.SeriesID).Scan(&id); err != nil {
		return err
	}
	season.SeasonID = id
	return nil
}

func (mr *MediaRepository) CreateEpisode(ctx context.Context, episode *domain.Episode) error {
	query := `INSERT INTO episode (title, description, number, season_id, video_url) VALUES ($1, $2, $3, $4, $5) returning episode_id`
	var id int64
	if err := mr.db.QueryRowxContext(ctx, query, episode.Title, episode.Description, episode.Number, episode.SeasonID, episode.VideoURL).Scan(&id); err != nil {
		return err
	}
	episode.EpisodeID = id
	return nil
}
