package link

import "time"

type UpdateRequest struct {
	OriginalURL string     `json:"original-url"`
	CustomCode  string     `json:"custom-code"`
	ExpiresAt   *time.Time `json:"expires-at"`
}
