package access


import(
	// "github.com/hosod/drive_access/internal/pkg"
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

// Options is flags
type Options struct {
	verbose bool `short:"v" long:"verbose"`
	// UploadCmd Upload
	// DownloadCmd Download
}

// Upload is subcommand
type Upload struct {
	Local string `short:"l" long:"local" required:"yes" description:"upload file from this local path"`
	Drive string `short:"d" long:"drive" required:"yes" description:"upload file to this path"`
}

// Execute is Upload process
func (upcmd *Upload) Execute(args []string) error {
	// some exec
	fmt.Println(upcmd.Local)
	fmt.Println(upcmd.Drive)
	return nil
}

// Download is subcommand
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
func GetParser() (*flags.Parser) {
	var opts Options
	var parser = flags.NewParser(&opts, flags.Default)

	var upcmd Upload
	var downcmd Download

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

	return parser
}