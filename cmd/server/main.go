package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
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

	api.FindOrCreatePersonHandler = operations.FindOrCreatePersonHandlerFunc(
		func(params operations.FindOrCreatePersonParams) middleware.Responder {
			name := swag.StringValue(params.LastName)

			greeting := fmt.Sprintf("http://rialto.stanford.org/person/%s", name)
			return operations.NewFindOrCreatePersonOK().WithPayload(greeting)
		})

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
