package link

import (
	"context"
	"encoding/json"
	dto "linkpulse/internal/dto/link"
	"testing"
)

func TestUpdateLink(t *testing.T) {
	for _, tc := range []struct {
		name, body string
		valid      bool
	}{
		{"empty", `{}`, false},
		{"url only", `{"original-url":"https://example.com"}`, true},
		{"clear expiry", `{"expires-at":null}`, true},
		{"future expiry", `{"expires-at":"2099-01-01T00:00:00Z"}`, true},
		{"past expiry", `{"expires-at":"2000-01-01T00:00:00Z"}`, false},
		{"null url", `{"original-url":null}`, false},
		{"empty code", `{"custom-code":""}`, false},
		{"null code", `{"custom-code":null}`, false},
		{"valid code", `{"custom-code":"my-code"}`, true},
		{"unsafe url", `{"original-url":"javascript:alert(1)"}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var req dto.UpdateRequest
			if err := json.Unmarshal([]byte(tc.body), &req); err != nil {
				t.Fatal(err)
			}
			err := New().UpdateLink(context.Background(), req)
			if (err == nil) != tc.valid {
				t.Fatalf("valid=%v, error=%v", tc.valid, err)
			}
			if tc.name == "clear expiry" && (!req.ExpiresAt.Set || !req.ExpiresAt.Null) {
				t.Fatal("null expiration was not preserved")
			}
			if tc.name == "url only" && (req.ExpiresAt.Set || req.CustomCode.Set) {
				t.Fatal("omitted fields marked as present")
			}
		})
	}
}
