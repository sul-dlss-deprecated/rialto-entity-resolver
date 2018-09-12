package handlers

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

// BuildAPI create new service API
func BuildAPI(registry *runtime.Registry) *operations.RialtoEntityResolverAPI {
	api := operations.NewRialtoEntityResolverAPI(swaggerSpec())

	api.FindOrCreatePersonHandler = NewFindOrCreatePerson(registry)
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
