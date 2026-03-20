package domain

type Media struct {
	MediaID      int64  `json:"mediaId" db:"media_id"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	TotalRatings int64  `json:"totalRatings" db:"total_ratings"`
	SumRatings   int64  `json:"sumRatings" db:"sum_ratings"`
	Tags         []Tag  `json:"tags" db:"tags"`
	Trailer      string `json:"trailer" db:"trailer"`
	Thumbnail    string `json:"thumbnail" db:"thumbnail"`
}
type Movie struct {
	Media
	VideoURL string `json:"videoURL" db:"video_url"`
}

type Series struct {
	Media
}

type Season struct {
	SeasonID int64  `json:"seasonId" db:"season_id"`
	Title    string `json:"title" db:"title"`
	Number   int64  `json:"number" db:"number"`
	SeriesID int64  `json:"seriesId" db:"series_id"`
}

type Episode struct {
	EpisodeID   int64  `json:"episodeId" db:"episode_id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	VideoURL    string `json:"videoURL" db:"video_url"`
	Number      int64  `json:"number" db:"number"`
	SeasonID    int64  `json:"seasonId" db:"season_id"`
}

type Tag string

const (
	ACTION      Tag = "action"
	ADVENTURE   Tag = "adventure"
	ANIMATION   Tag = "animation"
	BIOGRAPHY   Tag = "biography"
	COMEDY      Tag = "comedy"
	CRIME       Tag = "crime"
	DOCUMENTARY Tag = "documentary"
	DRAMA       Tag = "drama"
	FAMILY      Tag = "family"
	FANTASY     Tag = "fantasy"
	HORROR      Tag = "horror"
	MUSIC       Tag = "music"
	MYSTERY     Tag = "mystery"
	ROMANCE     Tag = "romance"
	SCI_FI      Tag = "sci_fi"
	THRILLER    Tag = "thriller"
	WAR         Tag = "war"
	WESTERN     Tag = "western"
)
