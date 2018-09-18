package main

import (
	"log"
	"os"
	"strconv"

	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi"
	"github.com/sul-dlss-labs/rialto-entity-resolver/handlers"
	"github.com/sul-dlss-labs/rialto-entity-resolver/repository"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

// var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	server := createServer()
	defer server.Shutdown()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer() *restapi.Server {
	repo := repository.BuildRepository()
	registry := runtime.NewRegistry(repo)
	api := handlers.BuildAPI(registry)
	server := restapi.NewServer(api)
	portFlag := os.Getenv("PORT")

	// Convert the ENV string to a number
	port, err := strconv.Atoi(portFlag)
	if err == nil {
		port = 3000
	}

	// A missing env variable will cause a "0" and not an error, so set a default
	if port == 0 {
		port = 3000
	}

	// set the port this service will be run on
	server.Port = port
	return server
}
