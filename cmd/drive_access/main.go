package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"errors"
	"io"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "../../configs/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
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


// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getService() (*drive.Service, error) {
	b,err := ioutil.ReadFile("../../configs/credentials.json")
	if err!=nil {
		fmt.Printf("Unable to read credentials.json file. Err: %v\n", err)
	}
	// If modifying these scopes, delete your previously saved token.json.
	config,err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err!=nil {
		return nil,err
	}
	client := getClient(config)
	
	service,err := drive.New(client)
	if err!=nil {
		fmt.Printf("Cannot create the Google Drive service: %v\n", err)
		return nil,err
	}
	return service,nil
}

func createDir(service *drive.Service, name string, parentID string) (*drive.File, error) {
	d := &drive.File {
		Name: name,
		MimeType: "application/vnd.google-apps.folder",
		Parents: []string{parentID},
	}
	file,err := service.Files.Create(d).Do()
	if err!=nil {
		log.Println("Could not create dir: " + err.Error())
		return nil,err
	}

	return file,nil
}

func createFile(service *drive.Service, name string, mimeType string, content io.Reader, parentID string) (*drive.File, error) {
	f := &drive.File {
		Name: name,
		MimeType: mimeType,
		Parents: []string{parentID},
	}
	file,err := service.Files.Create(f).Media(content).Do()
	if err!=nil {
		log.Println("Could not create file: " + err.Error())
		return nil, err
	}
	return file,nil
}

func getFileList(srv *drive.Service) ([]*drive.File, error){
	r, err := srv.Files.List().PageSize(20).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	fmt.Println("Files:")
	if len(r.Files) == 0 {
		return nil, errors.New("no files found")
	} 
	return r.Files, nil
}

func main() {
	// OPen the File 
	fmt.Println("Please input filename you want to upload.")
	var filename string
	if _,err := fmt.Scan(&filename); err!=nil {
		log.Fatalf("Unable to read authorization code: %v\n", err)
	}
	f,err := os.Open(filename)
	if err!=nil {
		log.Fatalf("Cannot open file: %\n", err)
	}
	defer f.Close()

	// Get the Google Drive service
	service, err := getService()

	// Create the Dir
	dir,err := createDir(service, "My Folder", "root")
	if err!=nil {
		log.Fatalf("Could not create dir: %v\n", err)
	}

	// Create the file and upload its content
	file,err := createFile(service, filename, "image/png", f, dir.Id)
	if err!=nil {
		log.Fatalf("Could not create file: %v\n", err)
	}

	log.Printf("File '%s' is successfully uploaded in '%s' directory", file.Name, dir.Name)	
}
