package handlers

import (
	"log"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

// BuildAPI create new service API
func BuildAPI(registry *runtime.Registry) *operations.RialtoEntityResolverAPI {
	api := operations.NewRialtoEntityResolverAPI(swaggerSpec())
	api.Logger = log.Printf
	// Applies when the "x-token" header is set
	api.KeyAuth = func(token string) (interface{}, error) {
		if token != os.Getenv("API_KEY") {
			api.Logger("Access attempt with incorrect api key auth: %s", token)
			return nil, errors.New(401, "incorrect api key auth")
		}
		// go-swagger will give a 401 if we don't return something.
		return "User", nil
	}

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
