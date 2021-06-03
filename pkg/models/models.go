package models

type Url struct {
	Url string `json:"url"`
}

// Spotify Struct

type AccessTokenResponse struct {
	SpotifyAccessToken string `json:"access_token"`
	TokenType          string `json:"token_type"`
	ExpiresIn          int    `json:"expires_in"`
	RefreshToken       string `json:"refresh_token"`
	Scope              string `json:"scope"`
}

type SpotifyProfile struct {
	Country      string `json:"country"`
	DisplayName  string `json:"display_name"`
	Email        string `json:"email"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		URL    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Product string `json:"product"`
	Type    string `json:"type"`
	URI     string `json:"uri"`
}
