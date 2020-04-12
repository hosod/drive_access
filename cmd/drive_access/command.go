package main

import(
	// "fmt"
	"log"
	"os"
	"path/filepath"

	// "golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func upload(service *drive.Service, localPath string, drivePath string) {
	files,err := filepath.Glob(localPath)
	if err!=nil {
		panic(err)
	}

	for _,file := range files {
		go func(file string) {
			f,err := os.Open(file)
			if err!=nil {
				log.Fatalf("Cannot open file: %\n", err)
			}
			defer f.Close()

			
		}(file)
		
	}
}