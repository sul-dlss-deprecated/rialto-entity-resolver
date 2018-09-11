package handlers

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
)

// BuildAPI create new service API
func BuildAPI() *operations.RialtoEntityResolverAPI {
	api := operations.NewRialtoEntityResolverAPI(swaggerSpec())

	api.FindOrCreatePersonHandler = NewFindOrCreatePerson()
	api.HealthCheckHandler = NewHealthCheck()

	return api
}

func swaggerSpec() *loads.Document {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	return swaggerSpec
}
