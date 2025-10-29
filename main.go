package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig *oauth2.Config

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .evn file: %s", err)
	}

	cliendID := os.Getenv("CLIENT_ID")
	cliend_secret := os.Getenv("CLIENT_SECRET")

	if cliendID == "" || cliend_secret == "" {
		log.Fatal("Error: CLIENT_ID or CLIENT_SECRET is not set in .env file")
	}

	googleOauthConfig = &oauth2.Config{
		ClientID:     cliendID,
		ClientSecret: cliend_secret,
		RedirectURL:  "http://localhost:8080/callback",
		Endpoint:     google.Endpoint,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
	}
	log.Println("OAuth Config setup complete.")

}
