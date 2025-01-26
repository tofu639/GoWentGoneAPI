// models.go
package models

// User model
type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Posts []Post `json:"posts"` // One-to-many relationship
}

// Post model
type Post struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"` // Foreign key
}
