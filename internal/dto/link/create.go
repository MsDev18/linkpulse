package link

import "time"

type CreateRequest struct {
	OriginalURL string     `json:"original-url" binding:"required"`
	CustomCode  string     `json:"custom-code"`
	ExpiresAt   *time.Time `json:"expires-at"`
}

type GetOneResponse struct {
	ID          uint       `json:"id"`
	OriginalURL string     `json:"original-url"`
	Code        string     `json:"code"`
	ShortURL    string     `json:"short-url"`
	ExpiresAt   *time.Time `json:"expires-at"`
	CreatedAt   time.Time  `json:"created-at"`
}
