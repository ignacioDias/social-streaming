package database

import "github.com/jmoiron/sqlx"

type Database struct {
	DB          *sqlx.DB
	UserRepo    *UserRepository
	SessionRepo *SessionRepository
	MediaRepo   *MediaRepository
	CastRepo    *CastRepository
	CommentRepo *CommentRepository
	ReviewRepo  *ReviewRepository
	FollowRepo  *FollowRepository
}

var tagsType string = `CREATE TYPE tag AS ENUM (
	'action',
	'adventure',
	'animation',
	'biography',
	'comedy',
	'crime',
	'documentary',
	'drama',
	'family',
	'fantasy',
	'horror',
	'music',
	'mystery',
	'romance',
	'sci_fi',
	'thriller',
	'war',
	'western'
)`

var usersTable string = `CREATE TABLE IF NOT EXISTS users (
	user_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	email TEXT NOT NULL UNIQUE,
	hashed_password TEXT NOT NULL,
	profile_picture TEXT NOT NULL DEFAULT '',
	banner_picture TEXT NOT NULL DEFAULT '',
	role SMALLINT NOT NULL CHECK (role IN (0, 1)),
	username TEXT NOT NULL UNIQUE,
	full_name TEXT NOT NULL
)`

var followsTable string = `CREATE TABLE IF NOT EXISTS follows (
	follow_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	follower_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	followed_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	UNIQUE (follower_id, followed_id),
	CHECK (follower_id <> followed_id)
)`

var sessionsTable string = `CREATE TABLE IF NOT EXISTS sessions (
	id TEXT PRIMARY KEY,
	user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	created_at TIMESTAMPTZ NOT NULL,
	expires_at TIMESTAMPTZ NOT NULL,
	CHECK (expires_at > created_at)
)`

var mediaTable string = `CREATE TABLE IF NOT EXISTS media (
	media_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL DEFAULT '',
	total_ratings BIGINT NOT NULL DEFAULT 0,
	sum_ratings BIGINT NOT NULL DEFAULT 0,
	tags tag[] NOT NULL DEFAULT '{}',
	CHECK (total_ratings >= 0),
	CHECK (sum_ratings >= 0)
)`

var moviesTable string = `CREATE TABLE IF NOT EXISTS movies (
	media_id BIGINT PRIMARY KEY REFERENCES media(media_id) ON DELETE CASCADE,
	video_url TEXT NOT NULL
)`

var seriesTable string = `CREATE TABLE IF NOT EXISTS series (
	media_id BIGINT PRIMARY KEY REFERENCES media(media_id) ON DELETE CASCADE
)`

var seasonsTable string = `CREATE TABLE IF NOT EXISTS seasons (
	season_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	title TEXT NOT NULL,
	number BIGINT NOT NULL,
	series_id BIGINT NOT NULL REFERENCES series(media_id) ON DELETE CASCADE,
	UNIQUE (series_id, number),
	CHECK (number > 0)
)`

var episodesTable string = `CREATE TABLE IF NOT EXISTS episodes (
	episode_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT NOT NULL DEFAULT '',
	video_url TEXT NOT NULL,
	number BIGINT NOT NULL,
	season_id BIGINT NOT NULL REFERENCES seasons(season_id) ON DELETE CASCADE,
	UNIQUE (season_id, number),
	CHECK (number > 0)
)`

var peopleTable string = `CREATE TABLE IF NOT EXISTS people (
	person_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name TEXT NOT NULL
)`

var castsTable string = `CREATE TABLE IF NOT EXISTS casts (
	cast_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	media_id BIGINT NOT NULL REFERENCES media(media_id) ON DELETE CASCADE,
	person_id BIGINT NOT NULL REFERENCES people(person_id) ON DELETE CASCADE,
	role SMALLINT NOT NULL CHECK (role IN (0, 1, 2, 3)),
	media_type SMALLINT NOT NULL CHECK (media_type IN (0, 1)),
	UNIQUE (media_id, person_id, role)
)`

var reviewsTable string = `CREATE TABLE IF NOT EXISTS reviews (
	review_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	score INT NOT NULL CHECK (score BETWEEN 0 AND 10),
	review TEXT NOT NULL DEFAULT '',
	target_id BIGINT NOT NULL,
	target_type SMALLINT NOT NULL CHECK (target_type IN (0, 1))
)`

var commentsTable string = `CREATE TABLE IF NOT EXISTS comments (
	comment_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_id BIGINT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	target_id BIGINT NOT NULL,
	content TEXT NOT NULL,
	target_type SMALLINT NOT NULL CHECK (target_type IN (0, 1, 2, 3))
)`

func (db *Database) Init() error {
	queries := []string{
		tagsType,
		usersTable,
		followsTable,
		sessionsTable,
		mediaTable,
		moviesTable,
		seriesTable,
		seasonsTable,
		episodesTable,
		peopleTable,
		castsTable,
		reviewsTable,
		commentsTable,
	}
	for _, q := range queries {
		if _, err := db.DB.Exec(q); err != nil {
			return err
		}
	}
	return nil
}
