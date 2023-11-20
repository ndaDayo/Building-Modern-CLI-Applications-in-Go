package main

import "net/http"

func main() {
	client := &http.Client{}

    cmds := []interfaces.Command{
        command.NewGetCommnad(client)
    }
}
