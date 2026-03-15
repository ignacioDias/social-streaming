package domain

type User struct {
	ID             int64  `db:"id" json:"id"`
	Email          string `json:"email" db:"email"`
	HashedPassword string `json:"-" db:"hashed_password"`
	ProfilePicture string `json:"profilePicture" db:"profile_picture"`
	Role           Role   `json:"role" db:"role"`
	Username       string `json:"username" db:"username"`
	FullName       string `json:"fullName" db:"full_name"`
}

type Role int

const (
	USER Role = iota
	ADMIN
)
