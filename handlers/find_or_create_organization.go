package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

// NewFindOrCreateOrganization will query Neptune for a organization record, or create a new one if no existing record can be found.
func NewFindOrCreateOrganization(registry *runtime.Registry) operations.FindOrCreateOrganizationHandler {
	return &findOrCreateOrganization{
		registry: registry,
	}
}

// findOrCreateOrganization handles a request for finding & returning an entry
type findOrCreateOrganization struct {
	registry *runtime.Registry
}

// Handle the retrieve resource request
func (d *findOrCreateOrganization) Handle(params operations.FindOrCreateOrganizationParams, principal interface{}) middleware.Responder {
	uri, err := d.registry.Repository.QueryForOrganizationByName(params.Name)

	if err != nil {
		log.Printf("%s", err)
		panic(err)
	}
	if uri != nil {
		return operations.NewFindOrCreateOrganizationOK().WithPayload(*uri)
	}

	uri, err = d.registry.Repository.CreateOrganization(params)
	if err != nil {
		panic(err)
	}
	return operations.NewFindOrCreateOrganizationOK().WithPayload(*uri)
}
