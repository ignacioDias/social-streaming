package domain

type Comment struct {
	CommentID  int64             `json:"commentId" db:"comment_id"`
	UserID     int64             `json:"userId" db:"user_id"`
	TargetID   int64             `json:"targetId" db:"target_id"`
	Content    string            `json:"content" db:"content"`
	TargetType CommentTargetType `json:"targetType" db:"target_type"`
}

type CommentTargetType int

const (
	POST CommentTargetType = iota
	MOVIE
	SERIES
	PROFILE
)
