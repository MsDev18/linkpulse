package entity

import "time"

type Link struct {
	ID          uint
	UserID      uint
	OriginalURL string
	Code        string
	ExpiresAt   *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
