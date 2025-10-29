package main

import (
	"fmt"
	"log"
	"net/http"
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
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)

	log.Println("Starting server on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}

}

func handleMain(w http.ResponseWriter, r *http.Request) {
	var html = `<html><body><a href="/login">Log in with Google</a></body></html>`
	fmt.Fprintf(w, html)
}

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("random-state-string")

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}
