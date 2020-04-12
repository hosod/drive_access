package main

import (
	"log"
	"os"
	// "fmt"
	// "strings"
	// "path/filepath"

	"github.com/hosod/drive_access/internal/pkg"

	flags "github.com/jessevdk/go-flags"
)


func main() {
	parser := access.GetParser()
	if _, err := parser.Parse(); err != nil {
		if fe, ok := err.(*flags.Error); ok && fe.Type == flags.ErrHelp {
			log.Fatalln(fe)
		}
		
		log.Print(err)
		os.Exit(1)
	}

	// path := "/root/hoge/fuga"
	// fmt.Println(strings.HasSuffix(path, "/"))
	
	// fmt.Println(filepath.Base(path))
	// fmt.Println(data)
	// fmt.Println(data[5:])
	// for d := range data[5:] {
	// 	fmt.Println(d)
	// }

	// srv,err := access.GetService()
	// id,err := access.ParseDrivePath(srv, path)
	// if err!=nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(id)

	// service, err := access.GetService()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// dirs, err := access.SearchFolder(service)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, dir := range dirs {
	// 	fmt.Println(dir.Name)
	// }
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
