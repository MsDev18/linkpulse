package auth

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh-token" binding:"required"`
}

type RefreshTokenResponse struct {
	Token Token `json:"token"`
}