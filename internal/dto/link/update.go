package link

import (
	"bytes"
	"encoding/json"
	"time"
)

// UpdateField distinguishes an omitted field from an explicit null.
type UpdateField[T any] struct {
	Set   bool
	Null  bool
	Value T
}

func (f *UpdateField[T]) UnmarshalJSON(data []byte) error {
	f.Set = true
	f.Null = bytes.Equal(bytes.TrimSpace(data), []byte("null"))
	if f.Null {
		var zero T
		f.Value = zero
		return nil
	}
	return json.Unmarshal(data, &f.Value)
}

type UpdateRequest struct {
	OriginalURL UpdateField[string]    `json:"original-url"`
	CustomCode  UpdateField[string]    `json:"custom-code"`
	ExpiresAt   UpdateField[time.Time] `json:"expires-at"`
}
