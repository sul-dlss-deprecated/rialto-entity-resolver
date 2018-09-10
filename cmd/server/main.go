package main

import (
	"flag"
	"log"

	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi"
	"github.com/sul-dlss-labs/rialto-entity-resolver/handlers"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	server := createServer()
	defer server.Shutdown()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer() *restapi.Server {
	api := handlers.BuildAPI()
	server := restapi.NewServer(api)

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag
	return server
}
