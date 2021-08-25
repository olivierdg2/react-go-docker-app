package main

import (
	server "github.com/olivierdg2/react-go-docker-app/go/pkg/cmd"
)

func main() {
	server.HandleRequests()
	defer server.Cli.Close()
}
