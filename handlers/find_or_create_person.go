package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss-labs/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss-labs/rialto-entity-resolver/runtime"
)

// NewFindOrCreatePerson will query Neptune for a person record, or create a new one if no existing record can be found.
func NewFindOrCreatePerson(registry *runtime.Registry) operations.FindOrCreatePersonHandler {
	return &findOrCreatePerson{
		registry: registry,
	}
}

// findOrCreatePerson handles a request for finding & returning an entry
type findOrCreatePerson struct {
	registry *runtime.Registry
}

// Handle the retrieve resource request
func (d *findOrCreatePerson) Handle(params operations.FindOrCreatePersonParams) middleware.Responder {
	if params.Orcid != nil {
		uri, err := d.registry.Repository.QueryForPersonByOrcid(*params.Orcid)
		if err != nil {
			panic(err)
		}
		if uri != nil {
			return operations.NewFindOrCreatePersonOK().WithPayload(*uri)
		}
	}

	if params.FirstName != nil && params.LastName != nil {
		uri, err := d.registry.Repository.QueryForPersonByName(*params.FirstName, *params.LastName)
		if err != nil {
			panic(err)
		}
		if uri != nil {
			return operations.NewFindOrCreatePersonOK().WithPayload(*uri)
		}
	}

	uri, err := d.registry.Repository.CreatePerson(params)
	if err != nil {
		panic(err)
	}
	return operations.NewFindOrCreatePersonOK().WithPayload(*uri)
}
