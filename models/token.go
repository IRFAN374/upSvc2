package models

type PayloadRequest struct {
	UserId int
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpired    string `json:"at_expired"`
	RtExpired    string `json:"rt_expired"`
}
