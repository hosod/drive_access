package main

import (
	"fmt"
	"log"
	// "os"

	// flags "github.com/jessevdk/go-flags"

	access "github.com/hosod/drive_access/internal/pkg"
)

// Options is
type Options struct {
	Command string `short:"c" long:"command" required:"yes" description:"[up, down, copy, delete, create]"`
	Path    string `short:"p" long:"path" required:"yes" description:"PATH/FROM/:PATH/TO/"`
}

var commands = []string{"up", "down", "copy", "delete", "create"}

var opts Options

func main() {
	service, err := access.GetService()
	if err != nil {
		fmt.Println(err)
	}
	// _,err = access.CreateDir(service, "hoge", "root")
	// if err!=nil {
	// 	log.Fatalln(err)
	// }
	dirs, err := access.SearchFolder(service)
	if err!=nil {
		log.Fatalln(err)
	}
	for _,dir := range dirs {
		fmt.Println(dir.Name)
	}
	// drivePath := "image/hoge"
	// file,err := access.ParseDriveDirPath(service, drivePath)
	// if err!=nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(file)
	// args,err := flags.Parse(&opts)
	// if err!=nil {
	// 	log.Fatalf("Err: %v", err)
	// }
	// fmt.Println(opts)
	// fmt.Println(args)

	// service,err := access.GetService()

	// switch opts.Command {
	// case "up":

	// case "down":

	// case "copy":

	// case "move":

	// case "delete":

	// case "create":

	// default:

	// }

	// // Open the File
	// fmt.Println("Please input filename you want to upload.")
	// var filename string
	// if _,err := fmt.Scan(&filename); err!=nil {
	// 	log.Fatalf("Unable to read authorization code: %v\n", err)
	// }
	// f,err := os.Open(filename)
	// if err!=nil {
	// 	log.Fatalf("Cannot open file: %\n", err)
	// }
	// defer f.Close()

	// // Get the Google Drive service
	// service, err := getService()

	// // Create the Dir
	// dir,err := createDir(service, "My Folder", "root")
	// if err!=nil {
	// 	log.Fatalf("Could not create dir: %v\n", err)
	// }

	// // Create the file and upload its content
	// file,err := createFile(service, filename, "image/png", f, dir.Id)
	// if err!=nil {
	// 	log.Fatalf("Could not create file: %v\n", err)
	// }

	// log.Printf("File '%s' is successfully uploaded in '%s' directory", file.Name, dir.Name)
}
