package main

import (
	"os"

	"github.com/spf13/pflag"

	"github.com/desperateofstruggle/Cloudgo/server"
)

const (
	// PORT - 8080 default port
	PORT string = "8080"
)

func main() {
	// port - get listened port
	port := os.Getenv("PORT")
	// if not exist then set to "8080"
	if len(port) == 0 {
		port = PORT
	}

	// bind the parameter p and parse
	pPort := pflag.StringP("port", "p", PORT, "PORT for http listening")
	pflag.Parse()

	// set the port
	if len(*pPort) != 0 {
		port = *pPort
	}

	// start the server
	server.NewServer(port)
}
