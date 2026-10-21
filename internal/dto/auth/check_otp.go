package auth

type CheckOtpRequest struct {
	PhoneNumber string `json:"phone-number" binding:"required"`
	Code        string `json:"code" binding:"required"`
}

type CheckOtpResponse struct {
	Token Token `json:"token"`
}

type Token struct {
	AccessToken  string `json:"access-token,omitempty"`
	RefreshToken string `json:"refresh-token,omitempty"`
}
