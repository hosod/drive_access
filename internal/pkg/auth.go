package access

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"encoding/json"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	
)

// GetClient retrives a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "../../configs/token.json"
	tok, err := TokenFromFile(tokFile)
	if err!=nil {
		tok = GetTokenFromWeb(config)
		SaveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// TokenFromFile retribes a token from a local file.
func TokenFromFile(file string) (*oauth2.Token, error) {
	f,err := os.Open(file)
	if err!=nil {
		return nil,err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// GetTokenFromWeb requests a token from the web, then returns the retrived token.
func GetTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// SaveToken saves a token to a file path.
func SaveToken(path string, token * oauth2.Token) {
	fmt.Printf("Saving credentials file to: %s\n", path)
	f,err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}



