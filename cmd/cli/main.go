package main

import (
	"net/http"

	"github.com/ndaDayo/Audio-Metadata-CLI/internal/interfaces"
)

func main() {
	client := &http.Client{}

    cmds := []interfaces.Command{
        command.NewGetCommand(client)
    }
}
