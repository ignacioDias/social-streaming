package domain

type Person struct {
	PersonID int64  `json:"personId" db:"person_id"`
	Name     string `json:"name" db:"name"`
}

type Cast struct {
	CastID     int64      `json:"castId" db:"cast_id"`
	MediaID    int64      `json:"mediaId" db:"media_id"`
	PersonID   int64      `json:"personId" db:"person_id"`
	PersonRole PersonRole `json:"role" db:"role"`
	MediaType  MediaType  `json:"mediaType" db:"media_type"`
}

type MediaType int

const (
	MOVIE_MEDIA MediaType = iota
	SERIES_MEDIA
)

type PersonRole int

const (
	ACTOR PersonRole = iota
	DIRECTOR
	PRODUCER
	WRITER
)
