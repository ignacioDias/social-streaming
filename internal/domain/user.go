package domain

type User struct {
	UserID         int64  `db:"user_id" json:"userId"`
	Email          string `json:"email" db:"email"`
	HashedPassword string `json:"-" db:"hashed_password"`
	ProfilePicture string `json:"profilePicture" db:"profile_picture"`
	BannerPicture  string `json:"bannerPicture" db:"banner_picture"`
	Role           Role   `json:"role" db:"role"`
	Username       string `json:"username" db:"username"`
	FullName       string `json:"fullName" db:"full_name"`
}

type Follow struct {
	FollowID   int64 `db:"follow_id" json:"followId"`
	FollowerID int64 `db:"follower_id" json:"followerId"`
	FollowedID int64 `db:"followed_id" json:"followedId"`
}

type Role int

const (
	USER Role = iota
	ADMIN
)
