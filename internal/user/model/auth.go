package model

// Token return to web
type Token struct {
	AccessToken string `json:"access_token"`
}

// Claims custom JWT claims
type Claims struct {
	UserID int `json:"user_id"`
}
