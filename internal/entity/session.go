package entity

import "time"

type Session struct {
	ID uint
	UserID uint
	ExpiresAt time.Time
	// nullable field 
	RevokeAt *time.Time
}