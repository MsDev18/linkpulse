package auth

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh-token"`
}

type RefreshTokenResponse struct {
	Token Token `json:"token"`
}