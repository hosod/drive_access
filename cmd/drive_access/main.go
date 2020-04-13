package main

import (
	"log"
	"os"
	
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
}
