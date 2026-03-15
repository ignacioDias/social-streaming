package domain

type Series struct {
	SeriesID     int64  `json:"seriesId" db:"series_id"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	SumRatings   int64  `json:"sumRatings" db:"sum_ratings"`
	TotalRatings int64  `json:"totalRatings" db:"total_ratings"`
	Tags         []Tag  `json:"tags" db:"tags"`
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
