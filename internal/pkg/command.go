package access

import (
	"fmt"
	"log"
	"strings"
	"path/filepath"

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
		return err
	}
	path := createcmd.Drive
	if strings.HasSuffix(path, "/") {
		path = filepath.Dir(path)
	}
	d,f := filepath.Split(path)

	id,err := ParseDrivePath(srv, d)
	if err!=nil {
		return err
	}
	
	_, err = CreateDir(srv, f, id)
	if err != nil {
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
	//some exec
	fmt.Println(downcmd.Local)
	fmt.Println(downcmd.Drive)
	return nil
}

// GetParser return parser
func GetParser() *flags.Parser {
	var opts Options
	var parser = flags.NewParser(&opts, flags.Default)

	var upcmd Upload
	var downcmd Download
	var createcmd Create

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

	return parser
}
