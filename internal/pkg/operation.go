package access

import (
	"errors"
	"fmt"
	"os"
	"io"
	"strings"
	"io/ioutil"
	"log"
	
	"path/filepath"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

// GetService create google drive service from credentials.json file.
func GetService() (*drive.Service, error) {
	exe,err := os.Executable()
	for i:=0;i<3;i++ {
		exe = filepath.Dir(exe)
	}
	credentials := filepath.Join(exe, "configs", "credentials.json")
	b, err := ioutil.ReadFile(credentials)
	if err != nil {
		fmt.Printf("Unable to read credentials.json file. Err: %v\n", err)
		return nil, err
	}
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, drive.DriveScope)
	if err != nil {
		return nil, err
	}
	client := GetClient(config)
	service, err := drive.New(client)
	if err != nil {
		fmt.Printf("Cannot create the Google Drive service. Err: %v\n", err)
		return nil, err
	}
	return service, nil
}

// CreateDir create dir under under the file that have given parentID.
func CreateDir(service *drive.Service, name string, parentID string) (*drive.File, error) {
	d := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
		Parents:  []string{parentID},
	}

	file, err := service.Files.Create(d).Do()
	if err != nil {
		log.Printf("Could not create dir. Err: %v\n", err)
		return nil, err
	}

	return file, nil
}

func CreateFile(service *drive.Service, name string, mimeType string, content io.Reader, parentID string) (*drive.File, error) {
	f := &drive.File{
		Name:     name,
		MimeType: mimeType,
		Parents:  []string{parentID},
	}
	file, err := service.Files.Create(f).Media(content).Do()
	if err != nil {
		log.Printf("Could not create file. Err: %v\n", err)
		return nil, err
	}
	return file, nil
}
// GetWholeFileList retribe list of all files
func GetWholeFileList(srv *drive.Service, parentID string) ([]*drive.File, error) {
	r, err := srv.Files.List().PageSize(20).
		Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		log.Printf("Unable to retrive files: %v\n", err)
		return nil, err
	}

	if len(r.Files) == 0 {
		return nil, errors.New("no files found")
	}
	return r.Files, nil
}
// SearchFolder search folder from dir of parentID and return id
func SearchFolder(srv *drive.Service, parentID string, folder string) (*drive.File, error) {
	r, err := srv.Files.List().
		Fields("nextPageToken, files(parents, id, name, mimeType)").
		// Do()
		// Q("'root' in parents").Do()
		Q(fmt.Sprintf("'%s' in parents and name='%s' and mimeType='application/vnd.google-apps.folder'",parentID, folder)).Do()
	if err != nil {
		return nil,err
	}
	switch len(r.Files) {
	case 0:
		return nil,fmt.Errorf("Unable to find dir: '%s'", folder)
	case 1:
		return r.Files[0],nil
	default:
		return nil,errors.New("many files are detected")
	}
}

func SearchFile(srv *drive.Service, parentID string, file string) (*drive.File, error) {
	r,err := srv.Files.List().
		Q(fmt.Sprintf("'%s' in parents and name='%s' and mimeType!='application/vnd.google-apps.folder'", parentID, file)).Do()
	if err!=nil {
		return nil,err
	}

	switch len(r.Files) {
	case 0:
		return nil,fmt.Errorf("Unable to find file: '%s'", file)
	case 1:
		return r.Files[0],nil
	default:
		return nil,errors.New("many files are detected")
	}
}

// ParseDrivePath retribe bottom element of given path
func ParseDrivePath(srv *drive.Service, path string) (string, error) {
	elements := strings.Split(path, "/")[2:]
	parentID := "root"
	for _,element := range elements {
		if element=="" {
			continue
		}
		f,err := SearchFolder(srv, parentID, element)
		if err!=nil {
			return "",err
		}
		parentID = f.Id
	}

	return parentID,nil
}

func DownloadFile(srv *drive.Service, fileID string, localpath string) (error) {
	// GET the file
	resp,err := srv.Files.Get(fileID).Download()
	if err!=nil {
		return err
	}
	defer resp.Body.Close()

	out,err := os.Create(localpath)
	if err!=nil {
		return err
	}
	defer out.Close()

	
	_,err = io.Copy(out, resp.Body)

	return err
}

// ParseDrivePath retribe bottom element of given path

