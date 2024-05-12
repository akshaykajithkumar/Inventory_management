package domain

// User represents a user in the system.
type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `gorm:"unique" json:"phone"`
}
