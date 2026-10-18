package entity

type User struct {
	ID uint
	// nullable field
	Name        *string
	PhoneNumber string
}
