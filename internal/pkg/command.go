package access

import (
	// "github.com/hosod/drive_access/internal/pkg"
	"fmt"
	"log"
	"os"
	"strings"
	"path/filepath"
	"text/tabwriter"

	flags "github.com/jessevdk/go-flags"

	// "golang.org/x/oauth2/google"
	// "google.golang.org/api/drive/v3"
)

// var srv *drive.Service

// Options is flags
type Options struct {
	verbose bool `short:"v" long:"verbose"`
	
}

// Upload is subcommand
type Upload struct {
	Local string `short:"l" long:"local" required:"yes" description:"upload file from this local path"`
	Drive string `short:"d" long:"drive" required:"yes" description:"upload file to this path"`
}

// Execute is Upload process
func (upcmd *Upload) Execute(args []string) error {
	// some exec
	localPaths, err := filepath.Glob(upcmd.Local)
	if err != nil {
		log.Println(err)
		return err
	}


	fmt.Println(localPaths)
	return nil
}
// Create is a subcommand for creating dir
type Create struct {
	Drive string `short:"d" long:"drive" required:"yes" decription:"create directory"`
}
// Execute is executed when "create" subcommand is called
func (createcmd *Create) Execute(args []string) error {
	srv, err := GetService()
	if err != nil {
		// log.Println(err)
		return err
	}
	path := createcmd.Drive
	if strings.HasSuffix(path, "/") {
		path = filepath.Dir(path)
	}
	d,f := filepath.Split(path)

	id,err := ParseDrivePath(srv, d)
	if err!=nil {
		// log.Println(err)
		return err
	}
	
	_, err = CreateDir(srv, f, id)
	if err != nil {
		// log.Println(err)
		return err
	}
	return nil
}

// Download is subcommand for download
type Download struct {
	Local string `short:"l" long:"local" required:"yes" description:"upload file to this local path"`
	Drive string `short:"d" long:"drive" required:"yes" description:"upload file from this drive path"`
}

// Execute is Download process
func (downcmd *Download) Execute(args []string) error {
	srv,err := GetService()
	if err!=nil {
		log.Println(err)
		return err
	}
	localPath := downcmd.Local
	drivePath := downcmd.Drive
	dir,file := filepath.Split(drivePath)

	dirID,err := ParseDrivePath(srv,dir)
	if err!=nil {
		log.Println(err)
		return err
	}
	
	driveFile,err := SearchFile(srv, dirID, file)
	if err!=nil {
		log.Println(err)
		return err
	}
	fileID := driveFile.Id

	err = DownloadFile(srv,fileID, filepath.Join(localPath, file))
	if err!=nil {
		log.Println(err)
		return err
	}
	return nil
}
// ListSegment is sub command for show list segment
type ListSegment struct {
	Path string `short:"p" long:"path" default:"/root" description:"show list segment of this path"`
}
// Execute is sclipt for execution listsegment sub command
func(ls *ListSegment) Execute(args []string) error {
	w := tabwriter.NewWriter(os.Stdout, 0,20,10,'\t',0)
	fmt.Fprintln(w,"Name\tType\tSize\tLink\t")

	srv,err := GetService()
	if err!=nil {
		log.Println(err)
		return err
	}

	parentID,err := ParseDrivePath(srv, ls.Path)
	if err!=nil {
		log.Println(err)
		return err
	}

	fileList,err := GetFileList(srv, parentID)
	if err!=nil {
		log.Println(err)
		return err
	}
	for _,file := range fileList {
		// fmt.Println(file.Name, file.MimeType, file.Size, file.WebViewLink)
		fmt.Fprintf(w, "%s\t%s\t%d\t%s\t\n", file.Name, file.MimeType, file.Size, file.WebViewLink)
	}
	w.Flush()
	return nil
}

// GetParser return parser
func GetParser() *flags.Parser {
	var opts Options
	var parser = flags.NewParser(&opts, flags.Default)

	var upcmd Upload
	var downcmd Download
	var createcmd Create
	var lscmd ListSegment

	parser.AddCommand(
		"upload",
		"uploadcmd",
		"",
		&upcmd,
	)

	parser.AddCommand(
		"download",
		"downloadcmd",
		"",
		&downcmd,
	)

	parser.AddCommand(
		"create",
		"createcmd",
		"",
		&createcmd,
	)

	parser.AddCommand(
		"ls",
		"listseg",
		"",
		&lscmd,
	)

	return parser
}
