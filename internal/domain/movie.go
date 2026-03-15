package domain

type Movie struct {
	MovieID      int64  `json:"movieId" db:"movie_id"`
	Title        string `json:"title" db:"title"`
	Description  string `json:"description" db:"description"`
	VideoURL     string `json:"videoURL" db:"video_url"`
	SumRatings   int64  `json:"sumRatings" db:"sum_ratings"`
	TotalRatings int64  `json:"totalRatings" db:"total_ratings"`
	Tags         []Tag  `json:"tags" db:"tags"`
}
