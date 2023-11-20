package command

import (
	"flag"
	"fmt"

	"github.com/ndaDayo/Audio-Metadata-CLI/internal/interfaces"
)

type UploadCommand struct {
	fs       *flag.FlagSet
	client   interfaces.Client
	filename string
}

func NewUploadCommand(client interfaces.Client) *UploadCommand {
	gc := &UploadCommand{
		fs:     flag.NewFlagSet("upload", flag.ContinueOnError),
		client: client,
	}

	gc.fs.StringVar(&gc.filename, "filename", "", "full path of filename to be uploaded")
	return gc
}

func (cmd *UploadCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *UploadCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: ./audiofile-cli upload -filename <filename>")
		return fmt.Errorf("missing flags")
	}

	return cmd.fs.Parse(flags)
}

func (cmd *UploadCommand) Run() error {
	return nil
}
