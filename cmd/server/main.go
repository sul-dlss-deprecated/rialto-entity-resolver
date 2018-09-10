package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewRialtoEntityResolverAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
