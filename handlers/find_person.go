package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/sul-dlss/rialto-entity-resolver/generated/models"
	"github.com/sul-dlss/rialto-entity-resolver/generated/restapi/operations"
	"github.com/sul-dlss/rialto-entity-resolver/runtime"
)

// NewFindPerson will query Neptune for a person record, or create a new one if no existing record can be found.
func NewFindPerson(registry *runtime.Registry) operations.FindPersonHandler {
	return &findPerson{
		registry: registry,
	}
}

// findPerson handles a request for finding & returning an entry
type findPerson struct {
	registry *runtime.Registry
}

// Handle the retrieve resource request
func (d *findPerson) Handle(params operations.FindPersonParams, principal interface{}) middleware.Responder {
	if params.Orcid != nil {
		uri, err := d.registry.Repository.QueryForPersonByOrcid(*params.Orcid)
		if err != nil {
			panic(err)
		}
		if uri != nil {
			return operations.NewFindPersonOK().WithPayload(*uri)
		}
	}

	if params.Sunetid != nil {
		uri, err := d.registry.Repository.QueryForPersonBySunetid(*params.Sunetid)
		if err != nil {
			panic(err)
		}
		if uri != nil {
			return operations.NewFindPersonOK().WithPayload(*uri)
		}
	}

	if params.FirstName != nil && params.LastName != nil {
		uri, err := d.registry.Repository.QueryForPersonByName(*params.FirstName, *params.LastName)
		if err != nil {
			panic(err)
		}
		if uri != nil {
			return operations.NewFindPersonOK().WithPayload(*uri)
		}
	}

	problem := &models.Error{
		Title:  "Not Found",
		Detail: "Unable to find an person matching the criteria you provided",
		Status: "404",
	}
	errors := []*models.Error{problem}

	return operations.NewFindPersonNotFound().WithPayload(&models.ErrorResponse{Errors: errors})
}
