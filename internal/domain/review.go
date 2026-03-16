package domain

type Review struct {
	ReviewID   int64     `json:"reviewId" db:"review_id"`
	UserID     int64     `json:"userId" db:"user_id"`
	Score      int       `json:"score" db:"score"`
	Review     string    `json:"review" db:"review"`
	TargetID   int64     `json:"targetId" db:"target_id"`
	TargetType MediaType `json:"targetType" db:"target_type"`
}
