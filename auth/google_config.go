package auth

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleAuthConfig interface {
	GetConfig() *oauth2.Config
}

type googleAuthConfig struct {
	config *oauth2.Config
}

func NewGoogleAuthConfig() GoogleAuthConfig {
	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return &googleAuthConfig{config: config}
}

func (g *googleAuthConfig) GetConfig() *oauth2.Config {
	return g.config
}
