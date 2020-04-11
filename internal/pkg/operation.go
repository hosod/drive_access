package access

import (
	"fmt"
	"log"
	"io/ioutil"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// GetService create google drive service from credentials.json file.
func GetService() (*drive.Service, error) {
	b, err := ioutil.ReadFile("../../configs/credentials.json")
	if err!=nil {
		fmt.Printf("Unable to read credentials.json file. Err: %v\n", err)
		return nil,err
	}
	// If modifying these scopes, delete your previously saved token.json.
	config,err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err!=nil {
		return nil,err
	}
	client := GetClient(config)
	service,err := drive.New(client)
	if err!=nil {
		fmt.Printf("Cannot create the Google Drive service. Err: %v\n", err)
		return nil,err
	}
	return service,nil
}

// CreateDir create dir under under the file that have given parentID.
func CreateDir(service *drive.Service, name string, parentID string) (*drive.File, error) {
	d := &drive.File {
		Name: name,
		MimeType: "application/vnd.google-apps.folder",
		Parents: []string{parentID},
	}

	file,err := service.Files.Create(d).Do()
	if err!=nil {
		log.Printf("Could not create dir. Err: %v\n", err)
		return nil, err
	}

	return file, nil
}

func CreateFile(service *drive.Service, name string, mimeType string, content io.Reader, parentID string) (*drive.File, error) {
	f := &drive.File {
		Name: name, 
		MimeType: mimeType,
		Parents: []string{parentID},
	}
	file,err := service.Files.Create(f).Media(content).Do()
	if err!=nil {
		log.Printf("Could not create file. Err: %v\n", err)
		return nil, err
	}
	return file, nil
}

// GetFileList retribe list of file 
// func GetFileList(srv *drive.Service) ([]*drive.File, error) {
// 	r,err := srv.Files.List().
// }

