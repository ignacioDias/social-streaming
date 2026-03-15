package domain

type Comment struct {
	UserID     int64      `json:"userId" db:"user_id"`
	TargetID   int64      `json:"targetId" db:"target_id"`
	Content    string     `json:"content" db:"content"`
	TargetType TargetType `json:"targetType" db:"target_type"`
}

type TargetType int

const (
	POST TargetType = iota
	MOVIE
	SERIES
	PROFILE
)
