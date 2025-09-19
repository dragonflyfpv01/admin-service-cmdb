package model

// UserList represents the users table structure
type UserList struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	UserID   string `json:"user_id" db:"user_id"`
	Role     string `json:"role" db:"role"`
}
