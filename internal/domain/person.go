package domain

type Person struct {
	PersonID int64  `json:"personId" db:"person_id"`
	Name     string `json:"name" db:"name"`
}

type MovieCast struct {
	MovieID    int64      `json:"movieId" db:"movie_id"`
	PersonID   int64      `json:"personId" db:"person_id"`
	PersonRole PersonRole `json:"role" db:"role"`
}

type SeriesCast struct {
	SeriesID   int64      `json:"seriesId" db:"series_id"`
	PersonID   int64      `json:"personId" db:"person_id"`
	PersonRole PersonRole `json:"role" db:"role"`
}

type PersonRole int

const (
	ACTOR PersonRole = iota
	DIRECTOR
	PRODUCTOR
	WRITER
)
