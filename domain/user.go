package domain

type UserID string

type User struct {
	ID          UserID
	Name        string
	Description string
	CreatedAt   int
	UpdatedAt   int
}
